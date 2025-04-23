package controller

import (
	dtorequest "server/internal/dto/request"
	"server/internal/usecase"
	"server/pkg/utils/ginutils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase usecase.UserUsecaseInterface
}

func NewUserController(
	UserUsecase usecase.UserUsecaseInterface,
) *UserController {
	return &UserController{
		UserUsecase: UserUsecase,
	}
}

func (sc *UserController) Get(ctx *gin.Context) {
	req := new(dtorequest.RequestGetUserByID)
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Error(err)
		return
	}
	res, err := sc.UserUsecase.GetUserByID(ctx, *req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ginutils.ResponseOK(ctx, res)
}
