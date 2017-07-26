package main

import (
	"fmt"
	"time"
)

func main() {
	// 	AsyncSample()
	// 	bench("SeriallProcess", SeriallProcess)
	// 	bench("ParallelProcess", ParallelProcess)
	// 	Bablesort()
	// 	err := raiseError()
	// 	log.Printf("%v\n", err)
	// 	log.Printf("%v", errors.Cause(err))
	// 	timeSample()
	// 	timeLocation()
	testInheritance()
}

func bench(label string, f func()) {
	start := time.Now().UnixNano()
	f()
	end := time.Now().UnixNano()
	fmt.Printf("%-15v : proctime %vms\n", label, (end-start)/1000000)
}
