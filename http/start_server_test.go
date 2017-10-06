package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"testing"
)

type queryParameter map[string][]string

type server struct {
	parameters chan queryParameter
}

func startServer() (*server, error) {
	server := new(server)
	params := make(queryParameter)
	server.parameters = make(chan queryParameter)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
		for key, val := range r.URL.Query() {
			params[key] = val
		}
		server.parameters <- params
	})

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		return nil, err
	}

	go http.Serve(l, handler)

	return server, nil
}

func main() {
	startServer()
}

func TestStartServer(t *testing.T) {
	_, err := startServer()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("should response 200, got %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(string(body)) != "Hello" {
		t.Errorf("want response body %v, got %v", "Hello", string(body))
	}
}

func TestStartServerQueryParam(t *testing.T) {
	s, err := startServer()
	if err != nil {
		t.Fatal(err)
	}

	if _, err := http.Get("http://localhost:8080?foo=bar&hoge=fuga"); err != nil {
		t.Fatal(err)
	}

	params := <-s.parameters
	foo, ok := params["foo"]
	if !ok || foo[0] != "bar" {
		t.Errorf("params foo want %v, got %v", "bar", foo)
	}
	hoge, ok := params["hoge"]
	if !ok || hoge[0] != "fuga" {
		t.Errorf("params hoge want %v, got %v", "fuga", hoge)
	}
}
