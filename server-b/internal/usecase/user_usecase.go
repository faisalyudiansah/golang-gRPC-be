package usecase

import (
	"context"
	dtorequest "server/internal/dto/request"
	dtoresponse "server/internal/dto/response"
	"server/internal/entity"
	clientgrpc "server/internal/grpc/client"
	"server/internal/repository"
	"server/pkg/database/transactor"
)

type UserUsecaseInterface interface {
	GetUserByID(ctx context.Context, req dtorequest.RequestGetUserByID) (*dtoresponse.ResponseUser, error)
}

type UserUsecaseImplementation struct {
	transactor     transactor.Transactor
	UserRepository repository.UserRepositoryInterface
	GRPCUserClient clientgrpc.UserClientGRPCInterface
}

func NewUserUsecaseImplementation(
	trx transactor.Transactor,
	UserRepo repository.UserRepositoryInterface,
	guc clientgrpc.UserClientGRPCInterface,
) *UserUsecaseImplementation {
	return &UserUsecaseImplementation{
		transactor:     trx,
		UserRepository: UserRepo,
		GRPCUserClient: guc,
	}
}

func (ir *UserUsecaseImplementation) GetUserByID(ctx context.Context, req dtorequest.RequestGetUserByID) (*dtoresponse.ResponseUser, error) {
	entityUser := new(entity.User)
	err := ir.transactor.Atomic(ctx, func(cForTx context.Context) error {
		resUserDB, err := ir.UserRepository.FindByID(cForTx, req.ID)
		if err != nil {
			return err
		}

		resGRPCUser, err := ir.GRPCUserClient.ClientGetUserByID(cForTx, req.ID)
		if err != nil {
			return err
		}

		entityUser.ID = resUserDB.ID
		entityUser.Role = resUserDB.Role
		entityUser.IsVerified = resUserDB.IsVerified
		entityUser.IsOAuth = resUserDB.IsOAuth
		entityUser.CreatedAt = resUserDB.CreatedAt
		entityUser.UpdatedAt = resUserDB.UpdatedAt
		entityUser.DeletedAt = resUserDB.DeletedAt
		entityUser.Email = resGRPCUser.Email
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &dtoresponse.ResponseUser{
		ID:         entityUser.ID,
		Role:       entityUser.Role,
		Email:      entityUser.Email,
		IsVerified: entityUser.IsVerified,
		IsOAuth:    entityUser.IsOAuth,
		CreatedAt:  entityUser.CreatedAt,
		UpdatedAt:  entityUser.UpdatedAt,
		DeletedAt:  entityUser.DeletedAt,
	}, nil
}
