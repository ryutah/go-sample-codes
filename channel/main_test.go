package main

import (
	"bufio"
	"bytes"
	"reflect"
	"testing"
)

func TestBufferChannel(t *testing.T) {
	type in struct {
		channelBuffers int
		loopCnt        int
	}
	cases := []struct {
		name string
		in   in
	}{
		{
			name: "buffer < message",
			in:   in{channelBuffers: 1, loopCnt: 10},
		},
		{
			name: "buffer = message",
			in:   in{channelBuffers: 10, loopCnt: 10},
		},
		{
			name: "buffer > message",
			in:   in{channelBuffers: 10, loopCnt: 5},
		},
		{
			name: "no buffers",
			in:   in{channelBuffers: 0, loopCnt: 5},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var w bytes.Buffer
			bufferChannel(&w, c.in.channelBuffers, c.in.loopCnt)

			s := bufio.NewScanner(&w)
			var lastLine string
			for s.Scan() {
				lastLine = s.Text()
			}
			if lastLine != "Finish" {
				t.Errorf(`last line should be "Finish", got: %q`, lastLine)
			}
		})
	}
}

func TestChannelRange(t *testing.T) {
	type (
		in struct {
			cnt int
		}
		out struct {
			ints []int
		}
	)
	cases := []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "test1",
			in: in{
				cnt: 5,
			},
			out: out{
				ints: []int{0, 1, 2, 3, 4},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := channelRange(c.in.cnt)
			if !reflect.DeepEqual(c.out.ints, got) {
				t.Errorf("\nwant: %v\n got: %v", c.out.ints, got)
			}
		})
	}
}
