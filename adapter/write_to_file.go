package main

import (
	"io/ioutil"
)

type FileWriter struct {
	file string
}

func NewFileWriter(file string) *FileWriter {
	return &FileWriter{file: file}
}

func (f *FileWriter) WriteMessage(msg string) error {
	return ioutil.WriteFile(f.file, []byte(msg), 770)
}
