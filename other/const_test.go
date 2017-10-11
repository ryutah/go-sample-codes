package main

import (
	"testing"
)

func TestConst(t *testing.T) {
	if FooConst != 1 {
		t.Errorf("want %v, got %v", 1, FooConst)
	}
	if BarConst != 2 {
		t.Errorf("want %v, got %v", 2, BarConst)
	}
	if HogeConst != 3 {
		t.Errorf("want %v, got %v", 3, HogeConst)
	}
}
