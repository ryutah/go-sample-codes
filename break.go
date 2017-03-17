package main

import (
	"fmt"
)

func Break() {
	fmt.Println("Switch Test please input 'h' : say hello, q' : quit")

loop:
	for {
		var str string
		fmt.Scanln(&str)
		switch str {
		case "h":
			fmt.Println("HELLO!!")
		case "q":
			fmt.Println("Quiet!!")
			break loop
		default:
			fmt.Println([]rune(str))
		}
	}

	fmt.Println("End")
}
