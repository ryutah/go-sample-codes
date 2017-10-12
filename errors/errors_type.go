package main

import (
	"github.com/pkg/errors"
)

type MyError struct {
	err error
}

func (m *MyError) Error() string {
	return m.err.Error()
}

type MyError2 struct{}

func (m *MyError2) Error() string {
	return "MyError2"
}

func errorsWithMessage() error {
	myErr2 := new(MyError2)
	err := errors.WithMessage(myErr2, "hogefuga")
	myErr := &MyError{err: err}
	return errors.WithMessage(myErr, "Failed")
}
