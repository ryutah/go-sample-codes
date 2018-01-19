package main

import (
	"testing"
)

func appendByIndex(size, loop int) {
	for i := 0; i < loop; i++ {
		s := make([]int, size)
		for j := 0; j < size; j++ {
			s[j] = j
		}
	}
}

func appendByAppend(size, loop int) {
	for i := 0; i < loop; i++ {
		s := make([]int, 0, size)
		for j := 0; j < size; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkAppendByIndex(b *testing.B) {
	const (
		size = 100000
		loop = 2000
	)
	for i := 0; i < b.N; i++ {
		b.Run("ByIndex", func(b *testing.B) {
			appendByIndex(size, loop)
		})
		b.Run("ByAppend", func(b *testing.B) {
			appendByAppend(size, loop)
		})
	}
}
