package workers

import (
	"context"
	"server/internal/gateway/provider"
	"server/pkg/logger"
)

func runGRPCWorker(ctx context.Context) {
	grpcServer := provider.ProvideGRPCDependency(cfg)
	if grpcServer == nil {
		logger.Log.Fatal("Failed to initialize gRPC provider")
	}

	go func() {
		if err := grpcServer.Start(); err != nil {
			logger.Log.Fatal("Failed to start gRPC server:", err)
		}
	}()

	<-ctx.Done()
	grpcServer.Shutdown()
}
