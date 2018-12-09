package gateway

import (
	"context"
	"fmt"
	"net/http"

	"github.com/colinking/budgeteer/backend/pkg/gen/userpb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

// Config specifies the expected environment variables
type Config struct {
	// Gateway server
	Port    int    `default:"9093" required:"false" split_words:"true"`
	CertDir string `default:"certs/" required:"false" split_words:"true"`

	// GRPC server
	GRPCPort int `default:"9091" required:"false" split_words:"true"`
}

// registerHTTPEndpoints returns a new gateway server which translates HTTP into gRPC.
func registerHTTPEndpoints(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{EmitDefaults: true}))

	for _, f := range []func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error{
		userpb.RegisterUserServiceHandler,
	} {
		if err := f(ctx, mux, conn); err != nil {
			return nil, err
		}
	}
	return mux, nil
}

func Run(ctx context.Context, c Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	addr := fmt.Sprintf("localhost:%d", c.Port)
	grpcAddr := fmt.Sprintf("localhost:%d", c.GRPCPort)

	tlsCertFilePath := c.CertDir + "localhost.crt"
	creds, err := credentials.NewClientTLSFromFile(tlsCertFilePath, "")
	if err != nil {
		return err
	}

	conn, err := grpc.DialContext(ctx,
		grpcAddr,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
	)
	if err != nil {
		return err
	}

	// TODO: standardize around chi
	mux := http.NewServeMux()
	gw, err := registerHTTPEndpoints(ctx, conn)
	if err != nil {
		return err
	}
	mux.Handle("/", gw)

	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		grpclog.Infof("Shutting down the http server")
		if err := s.Shutdown(context.Background()); err != nil {
			grpclog.Errorf("Failed to shutdown http server: %v", err)
		}
	}()

	grpclog.Infof("Starting listening at http://%s", addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		grpclog.Errorf("Failed to listen and serve: %v", err)
		return err
	}
	return nil
}
