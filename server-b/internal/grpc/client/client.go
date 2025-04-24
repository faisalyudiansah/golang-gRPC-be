package clientgrpc

import (
	"context"
	"fmt"
	pb "server/internal/grpc/proto/generate"
	"server/pkg/config"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClientGRPCInterface interface {
	ClientGetUserByID(ctx context.Context, id int64) (*pb.UserResponse, error)
	Close()
}

type UserClient struct {
	client pb.UserServiceClient
	conn   *grpc.ClientConn
	cfg    *config.Config

	initOnce sync.Once
	initErr  error
}

func NewUserClient(cfg *config.Config) *UserClient {
	return &UserClient{cfg: cfg}
}

func (c *UserClient) initClient() error {
	c.initOnce.Do(func() {
		address := fmt.Sprintf("%s:%s", c.cfg.GRPC.GRPCHost, "50052")

		conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			c.initErr = fmt.Errorf("failed to create gRPC client: %v", err)
			return
		}

		conn.Connect()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		state := conn.GetState()
		if ok := conn.WaitForStateChange(ctx, state); !ok {
			conn.Close()
			c.initErr = fmt.Errorf("timeout waiting for gRPC connection to be ready")
			return
		}

		c.conn = conn
		c.client = pb.NewUserServiceClient(conn)
	})
	return c.initErr
}

func (c *UserClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *UserClient) ClientGetUserByID(ctx context.Context, id int64) (*pb.UserResponse, error) {
	if err := c.initClient(); err != nil {
		return nil, err
	}

	req := &pb.GetUserRequest{Id: id}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := c.client.GetUserByID(ctxWithTimeout, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
