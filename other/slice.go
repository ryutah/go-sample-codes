package main

import (
	"fmt"
)

func appendToSlice(s *[]int, a int) {
	*s = append(*s, a)
}

func appendToSliceSample() {
	s := []int{1, 2}
	appendToSlice(&s, 4)
	fmt.Println(s)
}
