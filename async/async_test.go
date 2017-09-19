package main

import (
	"testing"
	"time"
)

func TestAsyncSample(t *testing.T) {
	ch := make(chan string)
	go asyncGet(ch)
	if s := <-ch; s == "" {
		t.Error("should not blank")
	}
}

func TestAsyncSample2(t *testing.T) {
	retval := asyncMap()
	val := <-retval
	if _, ok := val["hoge"]; !ok {
		t.Error("should get hoge")
	}
}

func asyncMap() chan map[string]string {
	ch := make(chan map[string]string)
	m := make(map[string]string)

	async := func() {
		time.Sleep(400 * time.Millisecond)
		m["hoge"] = "HOGEHOGE"
		ch <- m
	}

	go async()

	return ch
}
