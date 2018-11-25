package main

import (
	"context"
	"os"

	"github.com/colinking/budgeteer/backend/pkg/gateway"
	_ "github.com/jnewmano/grpc-json-proxy/codec"

	"google.golang.org/grpc/grpclog"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func loadGatewayConfigVars() (gateway.Config, error) {
	_ = godotenv.Overload()

	var c gateway.Config
	prefix := ""
	err := envconfig.Process(prefix, &c)

	return c, err
}

func main() {
	ctx := context.Background()

	// TODO: Standardize around a better logger
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stderr))
	c, err := loadGatewayConfigVars()
	if err != nil {
		grpclog.Fatalf("failed to load config: %v", err)
	}

	if err := gateway.Run(ctx, c); err != nil {
		grpclog.Fatal(err)
	}
}
