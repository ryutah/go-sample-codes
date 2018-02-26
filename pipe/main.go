package main

// ref http://blog.kaneshin.co/entry/2016/07/05/004601
// ref https://qiita.com/tanksuzuki/items/e712717675faf4efb07a

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	fmt.Printf("Receive Args : %v\n", os.Args)
	fmt.Println(terminal.IsTerminal(syscall.Stdin))
	if !terminal.IsTerminal(syscall.Stdin) {
		body, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Receive Pipeline : %v\n", string(body))
	}
}
