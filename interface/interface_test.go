package main

import (
	"testing"
)

func TestFoo(t *testing.T) {
	f := new(foo)
	if f.Bar() != "" {
		t.Error("want blank, got %v", f.Bar())
	}
}
