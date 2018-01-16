package main

const Foo = "foo"

func foo() string {
	return _foo(Foo)
}

func _foo(f string) string {
	return f
}
