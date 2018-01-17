package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const cnt = 600000
	fizzBuzz3(cnt)
}

func fizzBuzz(cnt int) {
	for i := 0; i < cnt; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func fizzBuzz2(cnt int) {
	buf := bufio.NewWriter(os.Stdout)
	for i := 0; i < cnt; i++ {
		if i%15 == 0 {
			fmt.Fprintln(buf, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Fprintln(buf, "Fizz")
		} else if i%5 == 0 {
			fmt.Fprintln(buf, "Buzz")
		} else {
			fmt.Fprintln(buf, i)
		}
	}
}

func fizzBuzz3(cnt int) {
	buf := bufio.NewWriter(os.Stdout)
	for i := 0; i < cnt; i++ {
		switch i % 15 {
		case 0:
			fmt.Fprintln(buf, "FizzBuzz")
		case 3:
			fmt.Fprintln(buf, "Buzz")
		case 5:
			fmt.Fprintln(buf, "Fizz")
		default:
			fmt.Fprintln(buf, i)
		}
	}
}
