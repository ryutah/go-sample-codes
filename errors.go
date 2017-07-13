package main

import (
	"os"

	"github.com/pkg/errors"
)

func raiseError() error {
	f, err := os.Open("nofile")
	if err != nil {
		return errors.WithMessage(err, "hogefuga")
	}
	defer f.Close()
	return nil
}
