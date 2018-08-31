package main

import (
	"testing"
)

func TestSelfCryption(t *testing.T) {
	type (
		in struct {
			caesar int
			enc    string
			dec    string
		}
		out struct {
			enc string
			dec string
		}
	)
	cases := []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "shuffle by 1",
			in: in{
				caesar: 1,
				enc:    "abcd",
				dec:    "bcde",
			},
			out: out{
				enc: "bcde",
				dec: "abcd",
			},
		},
		{
			name: "shuffle by 10",
			in: in{
				caesar: 10,
				enc:    "abcd",
				dec:    "klmn",
			},
			out: out{
				enc: "klmn",
				dec: "abcd",
			},
		},
		{
			name: "shuffle by 15",
			in: in{
				caesar: 15,
				enc:    "abcd",
				dec:    "pqrs",
			},
			out: out{
				enc: "pqrs",
				dec: "abcd",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cs := NewCaesar(c.in.caesar)
			if a, e := cs.Encryption(c.in.enc), c.out.enc; a != e {
				t.Errorf("Encryption\nwant: %v\n got: %v", e, a)
			}
			if a, e := cs.Decryption(c.in.dec), c.out.dec; a != e {
				t.Errorf("Decryption\nwant: %v\n got: %v", e, a)
			}
		})
	}
}
