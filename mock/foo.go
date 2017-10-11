package main

//go:generate mockgen -source=foo.go -package=main -destination=foo_mock.go
type Foo interface {
	Bar(string) string
}

func SampleFoo(foo Foo) string {
	return foo.Bar("In SampleFoo")
}
