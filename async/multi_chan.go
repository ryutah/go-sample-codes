package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type data struct {
	n     int
	delay int
}

func (d data) String() string {
	return fmt.Sprintf("index %v, delayed %vms", d.n, d.delay)
}

func multiChan(n int) []data {
	ch := make(chan data, n)
	wg := new(sync.WaitGroup)
	for i := 0; i < n; i++ {
		go sendIntAfterRandomDelay(i, ch, wg)
	}
	wg.Wait()
	close(ch)

	var results []data
	for c := range ch {
		log.Printf("Got... %v", c)
		results = append(results, c)
	}
	return results
}

func sendIntAfterRandomDelay(n int, c chan data, wg *sync.WaitGroup) {
	wg.Add(1)
	delay := rand.Intn(500)
	stop := time.Duration(delay) * time.Millisecond
	time.Sleep(stop)
	c <- data{n: n, delay: delay}
	wg.Done()
}
