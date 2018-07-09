package main

import "testing"

func TestRetCaller(t *testing.T) {
	file, line := retCaller()
	if file != "caller_test.go" {
		t.Errorf("want: %v, got %v", "caller_test.go", file)
	}
	if line != 6 {
		t.Errorf("want: %v, got %v", 6, line)
	}
}

func TestCallCaller(t *testing.T) {
	file, line := callCaller()
	if file != "main.go" {
		t.Errorf("want: %v, got %v", "main.go", file)
	}
	if line != 4 {
		t.Errorf("want: %v, got %v", 4, line)
	}
}
