package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func main() {
	client, err := newClient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if val, err := client.Get("FOO").Result(); err == redis.Nil {
		log.Println("Not set key")
	} else if err != nil {
		log.Printf("%#v", err)
	} else {
		fmt.Printf("Get Value: %v\n", val)
	}

	if err := client.Set("FOO", "BAR", 0).Err(); err != nil {
		log.Fatal(err)
	}

	val, err := client.Get("FOO").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Get Value: %v\n", val)

	if err := client.Del("FOO").Err(); err != nil {
		log.Fatal(err)
	}
}

func newClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}
	return client, nil
}
