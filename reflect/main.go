package main

import (
	"fmt"
	"reflect"
)

func main() {
	sliceSample()
}

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

func sliceSample() {
	var foo []int
	bar := []int{}
	hoge := make([]int, 10)
	sliceSampleTypeSlice(&foo)
	sliceSampleTypeSlice(&bar)
	sliceSampleTypeSlice(hoge)
}

func sliceSampleTypeSlice(v interface{}) {
	val := reflect.ValueOf(v)
	fmt.Println(val.CanSet())
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	fmt.Println(val.CanSet())
	fmt.Println(val.Type())
	fmt.Println(val.Type().Kind() == reflect.Slice)
	fmt.Println(val.Type().Elem())
	sType := reflect.SliceOf(val.Type().Elem())
	sVal := reflect.MakeSlice(sType, 10, 10)
	fmt.Println(sType, sVal.Index(0).CanSet())
	sVal.Index(0).Set(reflect.ValueOf(123))

	for i := 0; i < sVal.Len(); i++ {
		fmt.Println(sVal.Index(i).Interface())
	}
}
