package main

import (
	"fmt"
	"time"
)

type TimeFoo struct {
	Time time.Time
}

func CompareTime() {
	t := new(TimeFoo)
	compareTime := time.Time{}
	isSame := t.Time.Equal(compareTime)
	fmt.Println(isSame)
}
