package main

type WriterAdapter interface {
	WriteMessage(msg string) error
}
