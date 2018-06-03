package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world!!")
	})

	go http.ListenAndServe(":8080", server)

	fmt.Println("Start server on port 8080...\nPress [ENTER] to finish")
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
}
