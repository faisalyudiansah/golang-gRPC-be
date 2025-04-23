package handler

import (
	"context"
	dtorequest "server/internal/dto/request"
	pb "server/internal/grpc/proto/generate"
	"server/internal/usecase"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userUseCase usecase.UserUsecaseInterface
}

func NewUserHandler(userUseCase usecase.UserUsecaseInterface) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) GetUserByID(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	reqFormat := new(dtorequest.RequestGetUserByID)
	reqFormat.ID = req.Id
	user, err := h.userUseCase.GetUserByID(ctx, *reqFormat)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}
	return &pb.UserResponse{
		Id:        req.Id,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}
