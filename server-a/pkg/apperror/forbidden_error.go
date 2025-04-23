package apperror

import (
	"errors"

	"server/pkg/constant"
)

func NewForbiddenAccessError() *AppError {
	msg := constant.ForbiddenAccessErrorMessage

	err := errors.New(msg)

	return NewAppError(err, ForbiddenAccessErrorCode, msg)
}
