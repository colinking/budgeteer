package grpc

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/colinking/budgeteer/backend/pkg/auth"
	"github.com/colinking/budgeteer/backend/pkg/db"
	"github.com/colinking/budgeteer/backend/pkg/db/mysql"
	"github.com/colinking/budgeteer/backend/pkg/gen/userpb"
	"github.com/colinking/budgeteer/backend/pkg/handlers/user"
	"github.com/colinking/budgeteer/backend/pkg/plaid"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

// Config specifies the expected environment variables
type Config struct {
	// gRPC server
	Port    int    `default:"9091" required:"false" split_words:"true"`
	CertDir string `default:"certs/" required:"false" split_words:"true"`

	// MySQL
	DbPort     int    `default:"9092" required:"false" split_words:"true"`
	DbUsername string `default:"root" required:"false" split_words:"true"`
	DbPassword string `default:"password" required:"false" split_words:"true"`
	DbName     string `default:"moss" required:"false" split_words:"true"`

	// Plaid
	PlaidClientID  string `required:"true" split_words:"true"`
	PlaidPublicKey string `required:"true" split_words:"true"`
	PlaidSecret    string `required:"true" split_words:"true"`
	PlaidEnv       string `default:"sandbox" required:"false" split_words:"true"`
}

var (
	authKey       interface{}
	wrappedServer *grpcweb.WrappedGrpcServer
)

// registerEndpoints registers all API endpoints for a given gRPC server.
func registerEndpoints(server *grpc.Server, c Config, db db.Database) {
	plaidClient := plaid.New(c.PlaidClientID, c.PlaidPublicKey, c.PlaidSecret, c.PlaidEnv)
	userpb.RegisterUserServiceServer(server, user.New(&user.ServiceConfig{
		Database: db,
		Client:   plaidClient,
	}))
}

func isGatewayGrpcRequest(req *http.Request) bool {
	return req.Method == http.MethodPost && strings.HasPrefix(req.Header.Get("content-type"), "application/grpc")
}

func Run(c Config) error {
	// Open DB connection
	db, err := mysql.New(&mysql.Config{
		Port:         c.DbPort,
		DatabaseName: c.DbName,
		Username:     c.DbUsername,
		Password:     c.DbPassword,
	})
	if err != nil {
		return errors.Errorf("failed opening database: %v", err)
	}

	// Register gRPC endpoints
	grpcServer := grpc.NewServer()
	registerEndpoints(grpcServer, c, db)
	reflection.Register(grpcServer)

	// Wrap the grpc-web HTTP wrapper around the gRPC server
	wrappedServer = grpcweb.WrapServer(grpcServer)

	// Load our Auth0 public key to validate requests.
	jwks := auth.New()
	authKey, err = jwks.GetFirstValidationKey()
	if err != nil {
		return err
	}

	// Setup router middleware
	r := chi.NewRouter()
	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Heartbeat("/healthz"),
		middleware.Timeout(10*time.Second),
		validationMiddleware,
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

	err = http.ListenAndServeTLS(addr, tlsCertFilePath, tlsKeyFilePath, r)

	return errors.Errorf("failed starting server: %v", err)
}
