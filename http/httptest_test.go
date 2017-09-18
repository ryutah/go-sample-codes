package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMockServer(t *testing.T) {
	// TODO 何故かタイムアウトでテストが失敗する
	// 	ch := make(chan string)
	// ハンドラーを作成
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "sample")
		// XXX 多分ここでブロッキングしてるっぽい
		// 普通にプログラム起動してたら動くんだけど何故だろう
		// 		ch <- "Sample http"
	})

	// モックサーバを作成
	l, _ := net.Listen("tcp", ":8080")
	ts := httptest.Server{
		Listener: l,                              // ポートを設定したリスナーを設定
		Config:   &http.Server{Handler: handler}, // ハンドラーを設定
	}
	ts.Start()
	defer ts.Close()

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("should status code 200, got %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Trim(string(body), "\n") != "sample" {
		t.Errorf("want body %v, got %v", "sample", string(body))
	}
}

func TestHttpTest(t *testing.T) {
	sampleHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Test")
	})

	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()

	r, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(string(data)) != "Test" {
		t.Errorf("want %v, got %v", "Test", string(data))
	}
}
