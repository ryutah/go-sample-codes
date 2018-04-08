package main

import (
	"reflect"
	"testing"
)

func TestDefaultFoo(t *testing.T) {
	want := &Foo{
		optionFoo: 1,
		optionBar: "default bar",
	}
	got := NewFoo()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want : %+v, got %+v", want, got)
	}
}

func TestOptionFoo(t *testing.T) {
	want := &Foo{
		optionFoo: 2,
		optionBar: "option bar",
	}
	got := NewFoo(
		OptionBar("option bar"),
		OptionFoo(2),
	)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want : %+v, got %+v", want, got)
	}
}
