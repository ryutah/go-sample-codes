package main

import (
	"bytes"
	"testing"
)

func TestQueueSample(t *testing.T) {
	w := new(bytes.Buffer)
	queueSample(w)
	t.Logf("\n%v", w.String())
}
