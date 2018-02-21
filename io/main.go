package main

//go:generate mockgen -source main.go -destination main_mock.go -package main

type MyReader interface {
	Read(p []byte) (n int, err error)
}

type MyWriter interface {
	Write(p []byte) (n int, err error)
}
