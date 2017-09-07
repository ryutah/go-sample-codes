package main

import (
	"fmt"
	"math/rand"
	"time"
)

const SortCount = 50000

func Bablesort() {
	arr := createSortList()

	start := time.Now().UnixNano()
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	end := time.Now().UnixNano()

	fmt.Printf("Bablesort : proctime %vms\n", (end-start)/1000000)
}

func createSortList() []int {
	arr := make([]int, SortCount)
	for i := 0; i < SortCount; i++ {
		arr[i] = i
	}
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
