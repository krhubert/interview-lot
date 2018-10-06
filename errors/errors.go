package errors

import (
	"errors"
	"strings"
)

// Join combines multiple error into one sperated with new line.
func Join(errs []error) error {
	if len(errs) == 0 {
		return nil
	}

	errstrings := make([]string, len(errs))
	for i := range errs {
		errstrings[i] = errs[i].Error()
	}
	return errors.New(strings.Join(errstrings, "\n"))
}
