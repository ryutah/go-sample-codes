package main

import "testing"

func BenchmarkFib(b *testing.B) {
	b.Run("fib1", func(b *testing.B) {
		fib1(35)
	})
	b.Run("fib2", func(b *testing.B) {
		fib2(35)
	})
}
