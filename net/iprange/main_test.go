package main

import (
	"reflect"
	"testing"
)

func TestTestsIP(t *testing.T) {
	type (
		in struct {
			ipRange string
			testIP  string
		}
		out struct {
			ok  bool
			err error
		}
	)
	cases := []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "include(min)",
			in: in{
				ipRange: "132.17.2.56/28",
				testIP:  "132.17.2.48",
			},
			out: out{
				ok: true,
			},
		},
		{
			name: "include(max)",
			in: in{
				ipRange: "132.17.2.56/28",
				testIP:  "132.17.2.63",
			},
			out: out{
				ok: true,
			},
		},
		{
			name: "not include(min)",
			in: in{
				ipRange: "132.17.2.56/28",
				testIP:  "132.17.2.47",
			},
			out: out{
				ok: false,
			},
		},
		{
			name: "not include(max)",
			in: in{
				ipRange: "132.17.2.56/28",
				testIP:  "132.17.2.64",
			},
			out: out{
				ok: false,
			},
		},
		{
			name: "not CIDR",
			in: in{
				ipRange: "132.17.2.56",
				testIP:  "132.17.2.56",
			},
			out: out{
				ok: true,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ok, err := testsIP(c.in.ipRange, c.in.testIP)
			if a, e := err, c.out.err; !reflect.DeepEqual(a, e) {
				t.Errorf("return error\nwant: %#v\n got: %#v", e, a)
			}
			if a, e := ok, c.out.ok; a != e {
				t.Errorf("tests result\nwant: %v, got: %v", e, a)
			}
		})
	}
}
