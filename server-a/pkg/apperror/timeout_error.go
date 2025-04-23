package apperror

import (
	"errors"

	"server/pkg/constant"
)

func NewTimeoutError() *AppError {
	msg := constant.RequestTimeoutErrorMessage

	err := errors.New(msg)

	return NewAppError(err, RequestTimeoutErrorCode, msg)
}
