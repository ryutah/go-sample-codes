package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input")

	s.Scan()
	txt := s.Text()
	fmt.Println(txt)
}
