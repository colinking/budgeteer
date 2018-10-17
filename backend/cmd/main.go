package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/colinking/budgeteer/backend/pkg/db"
	"github.com/colinking/budgeteer/backend/pkg/db/dynamodb"
	"github.com/colinking/budgeteer/backend/pkg/handlers/plaid"
	"github.com/colinking/budgeteer/backend/pkg/handlers/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	plaid_proto "github.com/colinking/budgeteer/backend/pkg/proto/plaid"
	user_proto "github.com/colinking/budgeteer/backend/pkg/proto/user"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc/reflection"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
)

// Config specifies the expected environment variables
type Config struct {
	// gRPC server
	Port    int    `default:"9091" required:"false"`
	CertDir string `default:"certs/" required:"false"`

	// DynamoDB
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
func registerEndpoints(server *grpc.Server, c Config, db db.Database) {
	plaid_proto.RegisterPlaidServer(server, plaid.New(&plaid.ServiceConfig{
		ClientID: c.PlaidClientID,
		PublicKey: c.PlaidPublicKey,
		Secret: c.PlaidSecret,
		Env: c.PlaidEnv,
		Database: db,
	}))
	user_proto.RegisterUserServiceServer(server, user.New(&user.ServiceConfig{
		Database: db,
	}))
}

func startServer(c Config) {
	// Open DB connection
	db, err := dynamodb.New(&dynamodb.Config{
		Port: c.DbPort,
	})
	if err != nil {
		grpclog.Fatalf("failed opening database: %v", err)
	}

	// Register gRPC endpoints
	grpcServer := grpc.NewServer()
	registerEndpoints(grpcServer, c, db)
	reflection.Register(grpcServer)

	// Wrap the grpc-web HTTP wrapper around the gRPC server
	wrappedServer := grpcweb.WrapServer(grpcServer)
	grpcMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			grpclog.Infof("Handling chi request")
			if wrappedServer.IsAcceptableGrpcCorsRequest(req) || wrappedServer.IsGrpcWebRequest(req) {
				grpclog.Infof("-> appears to be a grpc request")
				wrappedServer.ServeHTTP(resp, req)
				return
			}
			grpclog.Infof("-> appears to be a non-grpc request")
			next.ServeHTTP(resp, req)
		})
	}

	// Setup router middleware
	r := chi.NewRouter()
	r.Use(
		chiMiddleware.Logger,
		chiMiddleware.Recoverer,
		chiMiddleware.Heartbeat("/healthz"),
		grpcMiddleware,
	)

	// Minimal handler to support above gRPC middleware.
	r.Get("/", func(resp http.ResponseWriter, req *http.Request) {
		io.WriteString(resp, "This server only supports gRPC endpoints.")
	})

	addr := fmt.Sprintf("localhost:%d", c.Port)
	grpclog.Infof("starting server listing on: https://%s", addr)

	// Start server listening on HTTPs
	tlsCertFilePath := c.CertDir + "localhost.crt"
	tlsKeyFilePath := c.CertDir + "localhost.key"
	if err := http.ListenAndServeTLS(addr, tlsCertFilePath, tlsKeyFilePath, r); err != nil {
		grpclog.Fatalf("failed starting server: %v", err)
	}
}

func main() {
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stderr))
	c, err := loadConfigVars()
	if err != nil {
		grpclog.Fatalf("failed to load config: %v", err)
	}

	startServer(c)
}
