package main

import "testing"

func TestParseTag(t *testing.T) {
	tg := tagSample{Foo: "foo"}
	got := parseTag(tg)
	if g := got.Get("Foo"); g != "foo : FOO, bar : BAR" {
		t.Errorf("want %q, got %q", "foo : FOO, bar : BAR", g)
	}
}
