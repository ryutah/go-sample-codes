package main

import (
	"errors"
	"reflect"
	"testing"
)

func setValue(v interface{}, fieldName string) error {
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return errors.New("v must be pointer")
	}
	val := reflect.ValueOf(v)
	val = reflect.Indirect(val)
	field := val.FieldByName(fieldName)
	field.Set(reflect.ValueOf("test"))
	return nil
}

func TestSetValue_AtPtrValue(t *testing.T) {
	type hoge struct {
		Hoge string
	}

	h := new(hoge)
	if err := setValue(h, "Hoge"); err != nil {
		t.Fatal(err)
	}

	if h.Hoge != "test" {
		t.Errorf("want : %v, got %v", "test", h.Hoge)
	}
}

func TestSetValue_AtValueObj(t *testing.T) {
	type hoge struct {
		Hoge string
	}

	var h hoge
	if err := setValue(h, "Hoge"); err == nil {
		t.Error("should be return error")
	}
}
