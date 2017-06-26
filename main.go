package main

import (
	"fmt"
	"time"
)

func main() {
	bench("SeriallProcess", SeriallProcess)
	bench("ParallelProcess", ParallelProcess)
}

func bench(label string, f func()) {
	start := time.Now().UnixNano()
	f()
	end := time.Now().UnixNano()
	fmt.Printf("%-15v : proctime %vms\n", label, (end-start)/1000000)
}
