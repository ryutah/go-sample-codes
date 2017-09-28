package main

import (
	"fmt"
	"reflect"
)

type typeSampler interface {
	Hoge()
}

type typeSampleImpl struct{}

func (t typeSampleImpl) Hoge() {
	fmt.Println("hoeghoge")
}

func typeSample() {
	t := typeSampleImpl{}
	t.Hoge()

	typeSampleIf()
}

func typeSampleIf() {
	modelType := reflect.TypeOf((*typeSampler)(nil)).Elem()
	fmt.Println(modelType)
}

type Fuga struct {
	Hoge string `json:"hoge"`
	Fuga int
	Foo  string `json:"foo"`
	Bar  float64
}

func fugaValues() map[string]interface{} {
	f := Fuga{
		Hoge: "hogehoge",
		Fuga: 111,
		Foo:  "FOO",
		Bar:  0.001,
	}
	val := reflect.ValueOf(f)
	typ := reflect.TypeOf(f)

	retVal := make(map[string]interface{})
	for i := 0; i < typ.NumField(); i++ {
		fld := typ.Field(i)
		name := fld.Tag.Get("json")
		if name == "" {
			name = fld.Name
		}
		field := val.Field(i)
		retVal[name] = field.Interface()
	}

	return retVal
}

func SetValuesByReflect(values map[string]interface{}) Fuga {
	dstFuga := &Fuga{}
	dstVal, dstTyp := reflect.ValueOf(dstFuga).Elem(), reflect.TypeOf(dstFuga).Elem()
	for key, val := range values {
		dfv := dstVal.FieldByName(key)
		dft, _ := dstTyp.FieldByName(key)
		vVal := reflect.ValueOf(val)
		dfv.Set(vVal.Convert(dft.Type))
	}
	return *dstFuga
}
