package main

import (
	"bytes"
	"testing"
)

func TestQueueSample(t *testing.T) {
	w := new(bytes.Buffer)
	queueSample(w)
	want := "1\n2\n"
	if got := w.String(); got != want {
		t.Errorf("want : %q, got : %q", want, got)
	}
}

func TestStackSample(t *testing.T) {
	w := new(bytes.Buffer)
	stackSample(w)
	want := "2\n1\n"
	if got := w.String(); got != want {
		t.Errorf("want : %q, got : %q", want, got)
	}
}
