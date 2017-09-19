package main

import (
	"fmt"
	"math/rand"
	"time"
)

func AsyncSample() {
	c := make(chan string)
	go asyncGet(c)
	f := func() {
		for {
			select {
			case s := <-c:
				fmt.Println(s)
				return
			case <-time.After(1000 * time.Millisecond):
				fmt.Println("Now loading...")
			}
		}
	}
	f()
}

func asyncGet(c chan string) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Float64()
	slpTime := r * 100
	time.Sleep(time.Duration(slpTime) * time.Millisecond)
	c <- fmt.Sprintf("Result after sleep %vms", slpTime)
}
