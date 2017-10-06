package main

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type Sample struct {
	ID   string `csv:"id"`
	Name string `csv:"name"`
	Age  int64  `csv:"age"`
}

func main() {
	f, err := os.Open("sample.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var payload []*Sample
	if err := gocsv.Unmarshal(f, &payload); err != nil {
		panic(err)
	}

	for _, p := range payload {
		log.Println(p)
	}
}
