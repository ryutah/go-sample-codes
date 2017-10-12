package main

import (
	"testing"
)

func TestMultiChan(t *testing.T) {
	gots := multiChan(10)
	if len(gots) != 10 {
		t.Error("want size %v, got %v", 10, len(gots))
	}
}
