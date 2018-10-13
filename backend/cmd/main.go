package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	budgeteer_proto "github.com/colinking/budgeteer/backend/pkg/proto/budgeteer"
	plaid_proto "github.com/colinking/budgeteer/backend/pkg/proto/plaid"
	"github.com/colinking/budgeteer/backend/pkg/services"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc/reflection"
)

// Config specifies the expected environment variables
type Config struct {
	// gRPC server
	Port    int    `default:"9091" required:"false"`
	CertDir string `default:"certs/" required:"false"`

	DbPort int `default:"9092" required:"false"`

	// Plaid
	PlaidClientID  string `required:"true" split_words:"true"`
	PlaidPublicKey string `required:"true" split_words:"true"`
	PlaidSecret    string `required:"true" split_words:"true"`
	PlaidEnv       string `required:"false" split_words:"true" default:"sandbox"`
}

func loadConfigVars() (Config, error) {
	_ = godotenv.Overload()

	var c Config
	prefix := ""
	err := envconfig.Process(prefix, &c)

	return c, err
}

// registerEndpoints registers all API endpoints for a given gRPC server.
func registerEndpoints(server *grpc.Server, c Config) {
	budgeteer_proto.RegisterBudgetServiceServer(server, services.NewBudgetService())
	plaid_proto.RegisterPlaidServer(server, services.New(c.PlaidClientID, c.PlaidPublicKey, c.PlaidSecret, c.PlaidEnv))
}

func startServer(c Config) {
	grpcServer := grpc.NewServer()
	registerEndpoints(grpcServer, c)
	reflection.Register(grpcServer)

	// TODO: how to pass in context with plaid object?
	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}

	addr := fmt.Sprintf("localhost:%d", c.Port)
	httpServer := http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(handler),
	}

	grpclog.Printf("Starting server listing on: https://%s", addr)

	tlsCertFilePath := c.CertDir + "localhost.crt"
	tlsKeyFilePath := c.CertDir + "localhost.key"
	if err := httpServer.ListenAndServeTLS(tlsCertFilePath, tlsKeyFilePath); err != nil {
		grpclog.Fatalf("failed starting server: %v", err)
	}
}

func main() {
	grpclog.SetLogger(log.New(os.Stdout, "cmd.main: ", log.LstdFlags))

	c, err := loadConfigVars()
	if err != nil {
		grpclog.Fatalf("failed to load config: %v", err)
	}

	startServer(c)
}
