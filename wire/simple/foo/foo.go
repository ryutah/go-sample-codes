package foo

import (
	"github.com/google/go-cloud/wire"
)

var SuperSet = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz)

type Foo struct {
	X int
}

func ProvideFoo() Foo {
	return Foo{X: 42}
}
