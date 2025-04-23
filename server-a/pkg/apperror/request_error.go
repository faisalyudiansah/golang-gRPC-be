package apperror

import (
	"errors"
	pkgconstant "server/pkg/constant"
)

func NewInvalidQueryLimitError() *AppError {
	msg := pkgconstant.InvalidQueryLimit
	err := errors.New(msg)
	return NewAppError(err, DefaultClientErrorCode, msg)
}

func NewInvalidQueryPageError() *AppError {
	msg := pkgconstant.InvalidQueryPage
	err := errors.New(msg)
	return NewAppError(err, DefaultClientErrorCode, msg)
}
