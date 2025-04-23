package provider

import (
	grpc "server/internal/grpc/server"
	"server/pkg/config"
)

func ProvideGRPCDependency(cfg *config.Config) *grpc.GRPCServer {
	gRPCProvideUserModule(cfg)
	return grpc.NewGRPCServer(cfg, userUseCase)
}
