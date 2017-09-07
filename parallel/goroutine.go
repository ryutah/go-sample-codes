package main

import (
	"fmt"
	"sync"
)

type Resource string

func InputOutput() {
	fmt.Println("Start polling input.")
	fmt.Println("Please type 'q' if you want quit program.")

	in, out := make(chan Resource), make(chan Resource)
	go func() {
		for {
			var str string
			fmt.Scanln(&str)
			in <- Resource(str)
			if str == "q" {
				close(in)
			}
		}
	}()

	go func() {
		for o := range out {
			fmt.Println(o)
		}
	}()

	lock := new(sync.WaitGroup)
	pollser(in, out, lock)
	lock.Wait()
}

func pollser(in, out chan Resource, lock *sync.WaitGroup) {
	lock.Add(1)
	for r := range in {
		s := fmt.Sprintf("INPUT %v", r)
		out <- Resource(s)
	}
	close(out)
	lock.Done()
}
