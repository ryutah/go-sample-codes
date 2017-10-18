package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	switch os.Args[1] {
	case "foo":
		foo()
	default:
		def()
	}
}

func def() {
	// オプションを追加
	str := flag.String("hoge", "", "usage of hoge")
	i := flag.Int("fuga", 0, "usage of int fuga")

	flag.Parse()

	// 使用方法を出力
	flag.Usage()

	fmt.Println(*str, *i, flag.Args())
}

func foo() {
	foo := flag.String("foo", "", "usage of foo")
	flag.Parse()
	fmt.Println("IN FOO", *foo, flag.Args())

	flag.Usage()

	fmt.Println(*foo)
}
