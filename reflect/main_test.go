package main

import (
	"reflect"
	"testing"
)

func TestFugaValues(t *testing.T) {
	want := map[string]interface{}{
		"hoge": "hogehoge",
		"Fuga": 111,
		"foo":  "FOO",
		"Bar":  0.001,
	}
	got := fugaValues()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestSetValuesByReflect(t *testing.T) {
	want := Fuga{
		Hoge: "hoge",
		Fuga: 1,
		Foo:  "foo",
		Bar:  0.01,
	}
	got := SetValuesByReflect(map[string]interface{}{
		"Hoge": want.Hoge,
		"Fuga": want.Fuga,
		"Foo":  want.Foo,
		"Bar":  want.Bar,
	})
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
