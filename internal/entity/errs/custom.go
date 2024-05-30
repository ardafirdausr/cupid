package errs

import "github.com/pkg/errors"

func NewErrProcessingData(customMessage ...string) error {
	return errors.Cause(ErrTypeProcessingData(messageOverride("Failed to process data", customMessage...)))
}

func NewErrInvalidData(customMessage ...string) error {
	return errors.Cause(ErrTypeInvalidData(messageOverride("Invalid Data", customMessage...)))
}

func NewErrUnauthorized(customMessage ...string) error {
	return errors.Cause(ErrTypeUnauthorized(messageOverride("Unauthorized", customMessage...)))
}

func NewErrInvalidToken(customMessage ...string) error {
	return errors.Cause(ErrTypeInvalidToken(messageOverride("Invalid Token", customMessage...)))
}

func NewErrExpiredToken(customMessage ...string) error {
	return errors.Cause(ErrTypeExpiredToken(messageOverride("Expired Token", customMessage...)))
}

func NewErrForbidden(customMessage ...string) error {
	return errors.Cause(ErrTypeForbidden(messageOverride("Forbidden", customMessage...)))
}

func NewErrNotFound(customMessage ...string) error {
	return errors.Cause(ErrTypeNotFound(messageOverride("Not Found", customMessage...)))
}

func NewErrUnprocessable(customMessage ...string) error {
	return errors.Cause(ErrTypeUnprocessable(messageOverride("Unprocessable", customMessage...)))
}

func NewErrInternal(customMessage ...string) error {
	return errors.Cause(ErrTypeInternal(messageOverride("Internal Server Error", customMessage...)))
}

func messageOverride(defaultMessage string, customMessage ...string) string {
	if len(customMessage) > 0 {
		return customMessage[0]
	}

	return defaultMessage
}
