package main

import (
	"reflect"
	"testing"
)

func TestTypeOfInterface(t *testing.T) {
	type hoge struct {
		A string
	}
	printType := func(v interface{}) {
		switch typ := v.(type) {
		case reflect.Value:
			t.Log("Type is reflect.Value")
		default:
			t.Logf("Type is %T", typ)
		}
	}
	h := hoge{}
	printType(h)
}
