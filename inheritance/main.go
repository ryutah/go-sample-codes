package main

import (
	"fmt"
)

func testInheritance() {
	a := &A{"HOGEHOGE"}
	b := &B{A: a, Fuga: "fugafug"}
	b.Call()
}

type A struct {
	Hoge string
}

func (a *A) Call() {
	fmt.Println(a.Hoge)
}

type B struct {
	*A
	Fuga string
}

func (b *B) Call() {
	b.A.Call()
	fmt.Println(b.Fuga)
}
