package main

import (
	"fmt"
	"testing"
)

func BenchmarkSort(b *testing.B) {
	type smap struct {
		key string
		f   func(int)
	}
	sorts := []smap{
		smap{key: "default", f: sortByDefault},
		smap{key: "interface", f: sortByInterface},
		smap{key: "manual", f: sortByManual},
	}
	for i := 10; i <= 1000000; i = i * 10 {
		for _, v := range sorts {
			b.Run(fmt.Sprintf("%d-%s", i, v.key), func(b *testing.B) {
				bench(b, b.N, i, v.f)
			})
		}
	}
}

func bench(b *testing.B, n, size int, f func(int)) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f(size)
	}
}
