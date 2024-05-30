package errs

import (
	"reflect"

	"github.com/pkg/errors"
)

func IsEqualType(err, comparator error) bool {
	if err == nil || comparator == nil {
		return false
	}

	if causer := errors.Cause(err); causer != nil {
		err = causer
	}

	return reflect.TypeOf(err) == reflect.TypeOf(comparator)
}
