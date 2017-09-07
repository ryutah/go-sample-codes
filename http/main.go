package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpClientSample() {
	client := new(http.Client)

	resp, err := client.Get("http://localhost:8080/fuga?fuga=f")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	resp2, err := client.Get("http://localhost:8080/fuga")
	if err != nil {
		fmt.Printf("ERR %T %v", err, err)
	}
	fmt.Println(resp2.Status)

	resp3, err := client.Get("http://localhost:8080/hogefuga")
	if err != nil {
		fmt.Printf("ERR %T %v", err, err)
	}
	fmt.Println(resp3.Status)
}
