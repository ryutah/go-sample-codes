package main

import (
	"flag"
	"fmt"
)

func main() {
	// オプションを追加
	str := flag.String("hoge", "", "usage of hoge")
	i := flag.Int("fuga", 0, "usage of int fuga")
	flag.Parse()

	// 使用方法を出力
	flag.Usage()

	fmt.Println(*str, *i, flag.Args())
}
