package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func execGracefulSample() {
	httpShutdown := make(chan struct{})

	server := http.Server{
		Addr:    ":8080",
		Handler: serveMux(httpShutdown),
	}
	defer server.Close()

	errChan := startGracefulServer(&server)
	sig, cancel := interruptShutdown()

	select {
	case err := <-errChan:
		log.Printf("Shutdown caused by error: %v\n", err)
	case s := <-sig:
		log.Printf("\nReceive signal: %v\n", s)
	case <-cancel:
		log.Println("Shutdown by stdin...")
	case <-httpShutdown:
		log.Println("Shutdown from http...")
	}

	ctx, can := context.WithTimeout(context.Background(), 5*time.Second)
	defer can()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

func serveMux(cancel chan<- struct{}) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hogehgoe")
		w.Write([]byte("hello!"))
	})
	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "FOO")
	})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		cancel <- struct{}{}
	})

	return mux
}

func startGracefulServer(s *http.Server) <-chan error {
	errChan := make(chan error, 1)
	go func() {
		log.Printf("Start server on %s\n", s.Addr)
		if err := s.ListenAndServe(); err != nil {
			errChan <- err
		}
	}()
	return errChan
}

func newServer(mux *http.ServeMux) *http.Server {
	return &http.Server{
		Handler: mux,
	}
}

func interruptShutdown() (sig <-chan os.Signal, cancel <-chan struct{}) {
	s := make(chan os.Signal, 1)
	c := make(chan struct{}, 1)

	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		bufio.NewScanner(os.Stdin).Scan()
		c <- struct{}{}
	}()
	return s, c
}
