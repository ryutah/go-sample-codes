package main

import (
	"testing"
	"time"
)

func limitGoroutine(n int, call func(int)) {
	c := make(chan bool, n)
	for i := 0; i < 1000; i++ {
		c <- true
		go func(j int) {
			defer func() {
				<-c
			}()
			time.Sleep(100 * time.Millisecond)
			call(j)
		}(i)
	}
}

func TestLimitGoroutine(t *testing.T) {
	limitGoroutine(100, func(i int) {
		t.Log(i)
	})
}
