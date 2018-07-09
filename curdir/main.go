package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	_, file, _, _ := runtime.Caller(0)
	fmt.Println(file)
	fmt.Println(os.Args)
	fmt.Println(os.Getwd())
}
