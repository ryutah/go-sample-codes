package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func raiseError() error {
	f, err := os.Open("nofile")
	if err != nil {
		err = errors.WithMessage(err, "hogefugahogefuga")
		return errors.WithMessage(err, "hogefuga")
	}
	defer f.Close()
	return nil
}

type Err1 struct {
	Msg string
}

func (e Err1) Error() string {
	return e.Msg
}

var Err1_1 = &Err1{
	Msg: "Err_1_1",
}

var Err1_2 = &Err1{
	Msg: "Err1_2",
}

type Err2 struct {
	Msg string
}

var Err2_1 = &Err2{
	Msg: "Err2_1",
}

func (e Err2) Error() string {
	return e.Msg
}

func errorCompareSample() {
	err := errCompareSampleInline()
	switch err {
	case Err1_1:
		fmt.Println("ERR1-1")
	case Err1_2:
		fmt.Println("ERR1-2")
	case Err2_1:
		fmt.Println("ERR2-1")
	default:
		fmt.Println("OTHER")
	}
}

func errCompareSampleInline() error {
	return Err2_1
}

func createNreError() error {
	return errors.New("sample error")
}
