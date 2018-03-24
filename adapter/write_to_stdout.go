package main

import (
	"fmt"
)

type StdOutWriter struct{}

func (s *StdOutWriter) WriteMessage(msg string) error {
	fmt.Println(msg)
	return nil
}
