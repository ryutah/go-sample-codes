package main

import (
	"strconv"
	"sync"
)

const ParallelSampleListCount = 30000000

func SeriallProcess() {
	arr := createParalellSampleArray()
	rtArr := make([]string, len(arr))
	for i, a := range arr {
		rtArr[i] = strconv.Itoa(a)
	}
}

func ParallelProcess() {
	arr := createParalellSampleArray()
	rtArr := make([]string, len(arr))

	wg := new(sync.WaitGroup)
	f := func(begin, end int) {
		for i := begin; i < end; i++ {
			rtArr[i] = strconv.Itoa(i)
		}
		wg.Done()
	}

	thCount := len(arr) / 10000
	if rem := len(arr) % 10000; rem != 0 {
		thCount++
	}

	for i := 0; i < thCount; i++ {
		from := i * 10000
		var to int
		if from+10000 >= len(arr) {
			to = len(arr)
		} else {
			to = from + 10000
		}
		wg.Add(1)
		go f(from, to)
	}
	wg.Wait()

	var emps []int
	for i, v := range rtArr {
		if v == "" {
			emps = append(emps, i)
		}
	}
}

func createParalellSampleArray() []int {
	arr := make([]int, ParallelSampleListCount)
	for i := 0; i < ParallelSampleListCount; i++ {
		arr[i] = i
	}
	return arr
}
