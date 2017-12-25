package main

import (
	"container/list"
	"fmt"
	"io"
)

func queueSample(w io.Writer) {
	l := list.New()
	// Push to queue
	l.PushBack(1)
	l.PushBack(2)

	// Dequeue from queue
	e := l.Remove(l.Front())
	fmt.Fprintln(w, e)
	e = l.Remove(l.Front())
	fmt.Fprintln(w, e)
}
