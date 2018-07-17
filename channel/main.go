package main

import (
	"fmt"
	"io"
	"sort"
	"sync"
)

func bufferChannel(w io.Writer, channelBuffers, loopCnt int) {
	var c chan string
	if channelBuffers > 0 {
		c = make(chan string, channelBuffers)
	} else {
		c = make(chan string)
	}
	for i := 0; i <= loopCnt; i++ {
		go func(index int) {
			if index != loopCnt {
				c <- fmt.Sprintf("Cnt: %v", index)
			} else {
				c <- "Finish"
			}
		}(i)
	}

	for {
		// It is not garantee to call certain time for goroutine count, because
		// the "Finish" message can be received before other goroutine message is
		// received.
		str := <-c
		fmt.Fprintln(w, str)
		if str == "Finish" {
			return
		}
	}
}

func channelRange(cnt int) []int {
	c := make(chan int)

	var (
		ret  []int
		lock sync.Mutex
	)
	go func() {
		for i := range c {
			lock.Lock()
			ret = append(ret, i)
			lock.Unlock()
		}
	}()

	for i := 0; i < cnt; i++ {
		c <- i
	}
	close(c)

	sort.Ints(ret)
	return ret
}
