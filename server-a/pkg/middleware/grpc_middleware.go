package middleware

import (
	"context"
	"server/pkg/logger"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

// GRPCLogger adalah middleware untuk logging request gRPC
func GRPCLogger() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()

		// Get client IP
		clientIP := "unknown"
		if p, ok := peer.FromContext(ctx); ok {
			clientIP = p.Addr.String()
		}

		// Get metadata
		md, _ := metadata.FromIncomingContext(ctx)

		// Process request with the handler
		resp, err := handler(ctx, req)

		// Log after processing
		params := map[string]any{
			"method":       info.FullMethod,
			"client_ip":    clientIP,
			"latency":      time.Since(start).String(),
			"metadata":     md,
			"request_type": req,
		}

		if err != nil {
			params["error"] = err.Error()
			logger.Log.WithFields(params).Error("incoming gRPC request")
		} else {
			logger.Log.WithFields(params).Info("incoming gRPC request")
		}

		return resp, err
	}
}

// GRPCStreamLogger adalah middleware untuk logging streaming gRPC
func GRPCStreamLogger() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		start := time.Now()

		// Get client IP
		ctx := ss.Context()
		clientIP := "unknown"
		if p, ok := peer.FromContext(ctx); ok {
			clientIP = p.Addr.String()
		}

		// Get metadata
		md, _ := metadata.FromIncomingContext(ctx)

		// Wrap the server stream to log received and sent messages
		wrapper := &serverStreamWrapper{
			ServerStream: ss,
			info:         info,
			clientIP:     clientIP,
		}

		// Process stream with handler
		err := handler(srv, wrapper)

		// Log after processing
		params := map[string]any{
			"method":           info.FullMethod,
			"client_ip":        clientIP,
			"latency":          time.Since(start).String(),
			"metadata":         md,
			"is_client_stream": info.IsClientStream,
			"is_server_stream": info.IsServerStream,
		}

		if err != nil {
			params["error"] = err.Error()
			logger.Log.WithFields(params).Error("incoming gRPC stream")
		} else {
			logger.Log.WithFields(params).Info("incoming gRPC stream")
		}

		return err
	}
}

// serverStreamWrapper adalah wrapper untuk grpc.ServerStream yang memungkinkan kita untuk
// mencatat operasi stream
type serverStreamWrapper struct {
	grpc.ServerStream
	info     *grpc.StreamServerInfo
	clientIP string
}

func (s *serverStreamWrapper) RecvMsg(m interface{}) error {
	err := s.ServerStream.RecvMsg(m)

	if err == nil {
		logger.Log.WithFields(map[string]any{
			"method":    s.info.FullMethod,
			"client_ip": s.clientIP,
			"recv_msg":  m,
		}).Debug("gRPC stream received message")
	}

	return err
}

func (s *serverStreamWrapper) SendMsg(m interface{}) error {
	err := s.ServerStream.SendMsg(m)

	if err == nil {
		logger.Log.WithFields(map[string]any{
			"method":    s.info.FullMethod,
			"client_ip": s.clientIP,
			"send_msg":  m,
		}).Debug("gRPC stream sent message")
	}

	return err
}
