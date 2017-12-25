package main

import (
	"container/list"
	"fmt"
	"io"
)

func queueSample(w io.Writer) {
	queue := list.New()
	// Push to queue
	queue.PushBack(1)
	queue.PushBack(2)

	// pop from queue
	e := queue.Remove(queue.Front())
	fmt.Fprintln(w, e)
	e = queue.Remove(queue.Front())
	fmt.Fprintln(w, e)
}

func stackSample(w io.Writer) {
	stack := list.New()

	// Push to stack
	stack.PushBack(1)
	stack.PushBack(2)

	// pop from stack
	e := stack.Remove(stack.Back())
	fmt.Fprintln(w, e)
	e = stack.Remove(stack.Back())
	fmt.Fprintln(w, e)
}
