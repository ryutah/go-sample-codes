// +build wireinject

package main

import (
	"context"

	"github.com/google/go-cloud/wire"
	"github.com/ryutah/go-sample-codes/wire/simple/foo"
)

func initializeBaz(ctx context.Context) (foo.Baz, error) {
	wire.Build(foo.SuperSet)
	return foo.Baz{}, nil
}
