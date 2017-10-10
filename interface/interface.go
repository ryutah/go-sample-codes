package main

type Foo interface {
	Bar() string
}

type foo struct {
	Foo
}
