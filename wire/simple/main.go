package main

import (
	"context"
	"log"
)

func main() {
	baz, err := initializeBaz(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", baz)
}
