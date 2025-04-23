package clientgrpc

import (
	"context"
	"fmt"
	"log"
	pb "server/internal/grpc/proto/generate"
	"server/pkg/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClientGRPCInterface interface {
	ClientGetUserByID(ctx context.Context, id int64) (*pb.UserResponse, error)
}

type UserClient struct {
	client pb.UserServiceClient
	conn   *grpc.ClientConn
}

func NewUserClient(cfg *config.Config) (*UserClient, error) {
	PortHost := fmt.Sprintf("%s:%s", cfg.GRPC.GRPCHost, cfg.GRPC.GRPCPort)
	conn, err := grpc.NewClient(PortHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	if err != nil {
		return nil, err
	}
	client := pb.NewUserServiceClient(conn)
	return &UserClient{
		client: client,
		conn:   conn,
	}, nil
}

func (c *UserClient) ClientGetUserByID(ctx context.Context, id int64) (*pb.UserResponse, error) {
	req := &pb.GetUserRequest{
		Id: id,
	}
	res, err := c.client.GetUserByID(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
