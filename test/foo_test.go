package main

import (
	"testing"
)

func Test_F(t *testing.T) {
	want := "FOO"
	if got := _foo("FOO"); got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}
