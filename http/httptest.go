package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
)

func mockServer() {
	ch := make(chan string)
	// ハンドラーを作成
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "sample")
		ch <- "Sample http"
	})

	// モックサーバを作成
	l, _ := net.Listen("tcp", ":8080")
	ts := httptest.Server{
		Listener: l,                              // ポートを設定したリスナーを設定
		Config:   &http.Server{Handler: handler}, // ハンドラーを設定
	}
	ts.Start()
	defer ts.Close()

	log.Println(<-ch)
}
