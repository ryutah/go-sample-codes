package main

import (
	"testing"
)

func TestPrintTopStack(t *testing.T) {
	err := createNreError()
	t.Logf("%+v", err)
}
