package main

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestGetCommandOutput(t *testing.T) {
	type in struct {
		cmd  string
		args []string
	}
	cases := []struct {
		name string
		in   in
		out  string
	}{
		{
			name: "echo",
			in: in{
				cmd:  "echo",
				args: []string{"hello"},
			},
			out: "hello",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out, err := getCommandOutput(c.in.cmd, c.in.args...)
			if err != nil {
				t.Fatal(err)
			}
			if got := strings.TrimSpace(out.stdout.String()); got != c.out {
				t.Errorf("\nwant: %q\n got: %q", got, c.out)
			}
		})
	}
}

func TestGetCommandOutput_Error(t *testing.T) {
	type in struct {
		cmd  string
		args []string
	}
	cases := []struct {
		name string
		in   in
		out  error
	}{
		{
			name: "echo",
			in: in{
				cmd:  "cat",
				args: []string{"nofile"},
			},
			out: errors.New("failed to exec command [cat nofile]: details: cat: nofile: No such file or directory"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := getCommandOutput(c.in.cmd, c.in.args...)
			if err == nil {
				t.Fatal("must be return error")
			}
			if !reflect.DeepEqual(c.out, err) {
				t.Errorf("\nwant: %v\n got: %v", c.out, err)
			}
		})
	}
}
