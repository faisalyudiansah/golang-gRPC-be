package controller

import (
	"server/internal/usecase"
	"server/pkg/utils/ginutils"

	"github.com/gin-gonic/gin"
)

type ExampleController struct {
	ExampleUsecase usecase.ExampleUsecaseInterface
}

func NewExampleController(
	ExampleUsecase usecase.ExampleUsecaseInterface,
) *ExampleController {
	return &ExampleController{
		ExampleUsecase: ExampleUsecase,
	}
}

func (sc *ExampleController) Get(ctx *gin.Context) {
	err := sc.ExampleUsecase.GetUsecase(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ginutils.ResponseOKPlain(ctx)
}
