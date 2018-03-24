package main

import (
	"log"
)

type Foo struct {
	name   string
	writer WriterAdapter
}

func (f Foo) writeFoo() error {
	return f.writer.WriteMessage(f.name)
}

func main() {
	f := &Foo{
		name:   "FOO",
		writer: new(StdOutWriter),
	}
	if err := f.writeFoo(); err != nil {
		log.Fatal(err)
	}

	f2 := &Foo{
		name:   "FOO2",
		writer: NewFileWriter("smple.txt"),
	}
	if err := f2.writeFoo(); err != nil {
		log.Fatal(err)
	}
}
