package main

import (
	"fmt"
	"time"
)

func timeSample() {
	now := time.Now()
	after := now.Add(time.Duration(3 * time.Hour)).Unix()
	fmt.Println(now)
	fmt.Println(time.Unix(after, 0), after)

	time.Sleep(1 * time.Second)
	after2 := time.Unix(now.Unix(), int64(10800*time.Second)).Unix()
	fmt.Println(now)
	fmt.Println(time.Unix(after2, 0), after)
}

type TimeFoo struct {
	Time time.Time
}

func CompareTime() {
	t := new(TimeFoo)
	compareTime := time.Time{}
	isSame := t.Time.Equal(compareTime)
	fmt.Println(isSame)
}

func timeLocation() {
	now := time.Now()
	t, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	jp, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	fmt.Println(now.In(t).Format("2006-01-02 15:04:05"))
	fmt.Println(now.In(jp).Format("2006-01-02 15:04:05"))
}
