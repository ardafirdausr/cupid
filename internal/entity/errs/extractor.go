package errs

import "github.com/pkg/errors"

func GetMessage(e error) string {
	errMessage := e.Error()
	if causer := errors.Cause(e); causer != nil {
		errMessage = causer.Error()
	}

	return errMessage
}

func GetCauserMessage(e error, defaultMessage string) string {
	if causer := errors.Cause(e); causer != nil {
		defaultMessage = causer.Error()
	}

	return defaultMessage
}
