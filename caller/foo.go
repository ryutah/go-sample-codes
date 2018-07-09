package main

import (
	"os"
	"runtime"
	"strings"
)

func retCaller() (file string, line int) {
	curdir, _ := os.Getwd()
	_, file, line, _ = runtime.Caller(1)
	file = strings.TrimPrefix(strings.Replace(file, curdir, "", 1), "/")
	return
}
