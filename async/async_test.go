package main

import (
	"testing"
)

func TestAsyncSample(t *testing.T) {
	ch := make(chan string)
	go asyncGet(ch)
	if s := <-ch; s == "" {
		t.Error("should not blank")
	}
}
