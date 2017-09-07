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
