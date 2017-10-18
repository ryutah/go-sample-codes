package main

import (
	"fmt"
	"reflect"
)

type tagSample struct {
	Foo string `foo:"FOO" bar:"BAR"`
}

type tagInfo map[string]string

func (t tagInfo) Get(key string) string {
	v := t[key]
	return v
}

func parseTag(v interface{}) tagInfo {
	vt := reflect.TypeOf(v)
	if vt.Kind() == reflect.Ptr {
		vt = vt.Elem()
	}
	tMap := make(tagInfo)
	for i := 0; i < vt.NumField(); i++ {
		f := vt.Field(i)
		ftv := f.Tag.Get("foo")
		btv := f.Tag.Get("bar")
		tMap[f.Name] = fmt.Sprintf("foo : %v, bar : %v", ftv, btv)
	}
	return tMap
}
