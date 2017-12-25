package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	hoge := flag.String("hoge", "", "hogehgoe")
	flag.Parse()
	fmt.Println(*hoge)
	fmt.Println(len(flag.Args()))

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("mockgen", "-source", filepath.Join(pwd, os.Getenv("GOFILE")))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(stdout)
	if err != nil {
		panic(err)
	}

	cmd2 := exec.Command("goimports")
	stdin, err := cmd2.StdinPipe()
	if err != nil {
		panic(err)
	}
	go func() {
		defer stdin.Close()
		writer := bufio.NewWriter(stdin)
		writer.WriteString(string(b))
		writer.Flush()
	}()

	stdout2, err := cmd2.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdout2.Close()

	if err := cmd2.Start(); err != nil {
		panic(err)
	}

	b2, err := ioutil.ReadAll(stdout2)
	if err != nil {
		panic(err)
	}
	fmt.Println("HHHHHHHHHHHHHHHHHHHH+++++++++++++")
	fmt.Println(string(b2))
}
