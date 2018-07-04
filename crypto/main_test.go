package main

import "testing"

func BenchmarkCreateKeyPairs(b *testing.B) {
	cases := []struct {
		name string
		in   int
	}{
		{
			name: "2048",
			in:   2048,
		},
	}

	for _, c := range cases {
		for i := 0; i < b.N; i++ {
			b.Run(c.name, func(b *testing.B) {
				if _, err := createKeyPairs(c.in); err != nil {
					b.Fatal(err)
				}
			})
		}
	}
}
