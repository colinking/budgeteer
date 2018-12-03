package main

import (
	"os"

	_ "github.com/jnewmano/grpc-json-proxy/codec"

	"github.com/colinking/budgeteer/backend/pkg/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func loadGRPCConfigVars() (grpc.Config, error) {
	_ = godotenv.Overload()

	var c grpc.Config
	prefix := ""
	err := envconfig.Process(prefix, &c)

	return c, err
}

func main() {
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stderr))
	c, err := loadGRPCConfigVars()
	if err != nil {
		grpclog.Fatalf("failed to load config: %v", err)
	}

	if err := grpc.Run(c); err != nil {
		grpclog.Fatal(err)
	}
}
