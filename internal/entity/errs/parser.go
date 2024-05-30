package errs

import (
	"net/http"

	"github.com/pkg/errors"
)

func GetHttpCode(err error) int {
	if causer := errors.Cause(err); causer != nil {
		err = causer
	}

	switch err.(type) {
	case ErrTypeInvalidData:
		return http.StatusBadRequest
	case ErrTypeUnauthorized, ErrTypeInvalidToken, ErrTypeExpiredToken:
		return http.StatusUnauthorized
	case ErrTypeForbidden:
		return http.StatusForbidden
	case ErrTypeNotFound:
		return http.StatusNotFound
	case ErrTypeProcessingData, ErrTypeUnprocessable:
		return http.StatusUnprocessableEntity
	case ErrTypeInternal:
		return http.StatusInternalServerError
	}

	return http.StatusInternalServerError
}
