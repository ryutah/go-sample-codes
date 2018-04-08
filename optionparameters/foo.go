package main

type Foo struct {
	optionFoo int
	optionBar string
}

func NewFoo(options ...fooOptionHnalder) *Foo {
	foo := &Foo{
		optionFoo: 1,
		optionBar: "default bar",
	}
	for _, option := range options {
		option(foo)
	}
	return foo
}

type fooOptionHnalder func(*Foo)

func OptionFoo(foo int) fooOptionHnalder {
	return func(f *Foo) {
		f.optionFoo = foo
	}
}

func OptionBar(bar string) fooOptionHnalder {
	return func(f *Foo) {
		f.optionBar = bar
	}
}
