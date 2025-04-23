package servergrpc

import (
	"fmt"
	"net"

	"server/internal/grpc/handler"
	pb "server/internal/grpc/proto/generate"
	"server/internal/usecase"
	"server/pkg/config"
	"server/pkg/logger"
	"server/pkg/middleware"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	cfg         *config.Config
	server      *grpc.Server
	userUseCase usecase.UserUsecaseInterface
}

func NewGRPCServer(
	cfg *config.Config,
	userUseCase usecase.UserUsecaseInterface,
) *GRPCServer {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middleware.GRPCLogger(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			middleware.GRPCStreamLogger(),
		)),
	)

	userHandler := handler.NewUserHandler(userUseCase)

	pb.RegisterUserServiceServer(server, userHandler)

	reflection.Register(server)

	return &GRPCServer{
		cfg:         cfg,
		server:      server,
		userUseCase: userUseCase,
	}
}

func (s *GRPCServer) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.cfg.GRPC.GRPCPort))
	if err != nil {
		return err
	}

	logger.Log.Info("Running gRPC server on port:", s.cfg.GRPC.GRPCPort)

	if err := s.server.Serve(lis); err != nil {
		logger.Log.Error("gRPC server error:", err)
		return err
	}

	return nil
}

func (s *GRPCServer) Shutdown() {
	logger.Log.Info("Shutting down gRPC server...")
	s.server.GracefulStop()
	logger.Log.Info("gRPC server stopped gracefully")
}
