package main

import (
	"testing"

	"github.com/pkg/errors"
)

func TestErrorsWithMessage(t *testing.T) {
	err := errorsWithMessage()
	if _, ok := errors.Cause(err).(*MyError); !ok {
		t.Errorf("want %v, got %T", "MyError", errors.Cause(err))
	}
	msg := err.Error()
	if msg != "Failed: hogefuga: MyError2" {
		t.Errorf("unexpected error message %v", msg)
	}
}
