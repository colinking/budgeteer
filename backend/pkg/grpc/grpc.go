package grpc

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/colinking/budgeteer/backend/pkg/auth"
	"github.com/colinking/budgeteer/backend/pkg/db"
	"github.com/colinking/budgeteer/backend/pkg/db/mysql"
	"github.com/colinking/budgeteer/backend/pkg/gen/plaidpb"
	"github.com/colinking/budgeteer/backend/pkg/gen/userpb"
	"github.com/colinking/budgeteer/backend/pkg/handlers/plaid"
	"github.com/colinking/budgeteer/backend/pkg/handlers/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
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

// registerEndpoints registers all API endpoints for a given gRPC server.
func registerEndpoints(server *grpc.Server, c Config, db db.Database) {
	plaidpb.RegisterPlaidServer(server, plaid.New(&plaid.ServiceConfig{
		ClientID:  c.PlaidClientID,
		PublicKey: c.PlaidPublicKey,
		Secret:    c.PlaidSecret,
		Env:       c.PlaidEnv,
		Database:  db,
	}))
	userpb.RegisterUserServiceServer(server, user.New(&user.ServiceConfig{
		Database: db,
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
	wrappedServer := grpcweb.WrapServer(grpcServer)
	grpcMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			grpclog.Infof("Received request with C-T: %s\n", req.Header.Get("content-type"))
			if wrappedServer.IsAcceptableGrpcCorsRequest(req) {
				grpclog.Infof("Incoming grpc cors-preflight request")
				wrappedServer.ServeHTTP(resp, req)
			} else if wrappedServer.IsGrpcWebRequest(req) || isGatewayGrpcRequest(req) {
				grpclog.Infof("Incoming grpc request")
				wrappedServer.ServeHTTP(resp, req)
			} else {
				grpclog.Infof("Incoming non-grpc request")
				next.ServeHTTP(resp, req)
			}
		})
	}

	// Load our Auth0 public key to validate requests.
	jwks := auth.New()
	key, err := jwks.GetFirstValidationKey()
	if err != nil {
		return err
	}

	// Add middleware to validate that all gRPC requests are authorized with Auth0.
	validationMiddle := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			// TODO: rework ordering to require all endpoints beyond a certain point to
			// be authenticated
			if wrappedServer.IsGrpcWebRequest(req) || isGatewayGrpcRequest(req) {
				// Perform auth validation only on gRPC requests.
				// TODO: perform extra validation on the iss/aud/exp fields.
				verifier := jwtauth.Verifier(jwtauth.New("RS256", nil, key))
				verifier(jwtauth.Authenticator(next)).ServeHTTP(resp, req)
			} else {
				next.ServeHTTP(resp, req)
			}
		})
	}

	// Setup router middleware
	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.Heartbeat("/healthz"),
		validationMiddle,
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
