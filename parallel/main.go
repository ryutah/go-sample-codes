package main

import (
	"fmt"
	"sync"
)

func ChannelSelect() {
	f := func(msg string, n int) (chan string, chan bool) {
		ch := make(chan string)
		done := make(chan bool)
		go func() {
			for i := 0; i < n; i++ {
				ch <- msg + " please!"
			}
			done <- true
		}()
		return ch, done
	}

	ch1, done1 := f("beer", 4)
	ch2, done2 := f("juice", 2)
	ch3, done3 := f("water", 1)

	f2 := func() {
		count := 0
		for {
			select {
			case msg := <-ch1:
				fmt.Println(msg)
			case msg := <-ch2:
				fmt.Println(msg)
			case msg := <-ch3:
				fmt.Println(msg)
			case <-done1:
				count++
			case <-done2:
				count++
			case <-done3:
				count++
			}

			if count > 2 {
				return
			}
		}
	}

	f2()
	fmt.Println("exit")
}

func ChannelSelect2() {
	f := func(msg string, n int) chan string {
		ch := make(chan string)
		go func() {
			for i := 0; i < n; i++ {
				ch <- msg + " please!"
			}
			close(ch)
		}()
		return ch
	}

	ch1 := f("beer", 4)
	ch2 := f("juice", 2)
	ch3 := f("water", 1)

	var wg sync.WaitGroup

	reportter := func(ch chan string) {
		defer wg.Done()
		for msg := range ch {
			fmt.Println(msg)
		}
	}

	wg.Add(3)
	go reportter(ch1)
	go reportter(ch2)
	go reportter(ch3)
	wg.Wait()

	fmt.Println("exit")
}
