package main

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkStringsBuilder(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		builder.WriteString(strings.Repeat("A", 1000))
	}
}

func BenchmarkBuffer(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		buf.WriteString(strings.Repeat("A", 1000))
	}
}
