package main

import (
	"log"
	"reflect"
)

type mergeSampleStruct struct {
	Foo    string
	Bar    int
	Hoge   string
	Fuga   string
	FooBar int64
}

func mergeStruct(src interface{}, compare interface{}, dst interface{}) {
	srcV, compV, dstV := reflect.ValueOf(src).Elem(), reflect.ValueOf(compare).Elem(), reflect.ValueOf(dst).Elem()
	srcT := reflect.TypeOf(src).Elem()
	for i := 0; i < srcT.NumField(); i++ {
		fn := srcT.Field(i).Name
		fv := srcV.FieldByName(fn)
		log.Printf("field : %v", fn)
		cfv, dfv := compV.FieldByName(fn), dstV.FieldByName(fn)

		set := fv
		switch fv.Type().Kind() {
		case reflect.String:
			cs := cfv.String()
			if cs != "" {
				set = cfv
			}
		case reflect.Int, reflect.Int32, reflect.Int64:
			ci := cfv.Int()
			if ci != 0 {
				set = cfv
			}
		case reflect.Ptr:
			if cfv.Interface() != nil {
				set = cfv
			}
		}
		dfv.Set(set)
	}
}
