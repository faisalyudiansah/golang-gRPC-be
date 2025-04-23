package apperror

import "server/pkg/constant"

func NewServerError(err error) *AppError {
	msg := constant.InternalServerErrorMessage

	return NewAppError(err, DefaultServerErrorCode, msg)
}
