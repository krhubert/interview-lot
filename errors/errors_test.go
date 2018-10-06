package errors

import (
	"errors"
	"testing"
)

func TestJoin(t *testing.T) {
	if err := Join([]error{}); err != nil {
		t.Fatalf("join errors failed - expected nil error")
	}

	errs := []error{errors.New("e1"), errors.New("e2")}
	if err := Join(errs); err.Error() != "e1\ne2" {
		t.Fatalf("join errors failed - want: %q, got: %q", "e1\ne2", err.Error())
	}
}
