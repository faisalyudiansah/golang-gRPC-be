package apperror

import (
	"errors"

	"server/pkg/constant"
)

func NewUnauthorizedError() *AppError {
	msg := constant.UnauthorizedErrorMessage

	err := errors.New(msg)

	return NewAppError(err, UnauthorizedErrorCode, msg)
}
