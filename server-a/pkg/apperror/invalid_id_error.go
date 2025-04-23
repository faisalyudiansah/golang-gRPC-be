package apperror

import (
	"errors"

	"server/pkg/constant"
)

func NewInvalidIdError() *AppError {
	msg := constant.InvalidIDErrorMessage

	err := errors.New(msg)

	return NewAppError(err, DefaultClientErrorCode, msg)
}
