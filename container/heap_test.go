package main

import (
	"reflect"
	"testing"
)

func TestMyHeapLen(t *testing.T) {
	m := &myHeap{l: []int{1, 2, 3}}
	if want, got := 3, m.Len(); want != got {
		t.Errorf("want : %v, got : %v", want, got)
	}
}

func TestMyHeapLess(t *testing.T) {
	m := &myHeap{l: []int{1, 2, 3}}
	if !m.Less(0, 1) {
		t.Error("shoud be true, got false")
	}
	if m.Less(1, 0) {
		t.Error("shoud be false, got true")
	}
}

func TestMySampleSwap(t *testing.T) {
	m := &myHeap{l: []int{1, 2, 3}}
	want := []int{2, 1, 3}
	m.Swap(0, 1)
	if !reflect.DeepEqual(m.l, want) {
		t.Errorf("want : %#v, got : %#v", want, m.l)
	}
}

func TestMyHeapPush(t *testing.T) {
	m := &myHeap{l: []int{1, 2, 3}}
	want := []int{1, 2, 3, 4}
	m.Push(4)
	if !reflect.DeepEqual(m.l, want) {
		t.Errorf("want : %#v, got : %#v", want, m.l)
	}
}

func TestMyHeapPop(t *testing.T) {
	m := &myHeap{l: []int{1, 2, 3}}
	wantList := []int{1, 2}
	if want, got := m.Pop(), 3; want != 3 {
		t.Errorf("want : %v, got : %v", want, got)
	}
	if !reflect.DeepEqual(m.l, wantList) {
		t.Errorf("want : %#v, got : %#v", wantList, m.l)
	}
}
