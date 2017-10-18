package main

import "testing"

func TestMergeStruct(t *testing.T) {
	s := &mergeSampleStruct{
		Foo:    "FOO",
		Bar:    1,
		Hoge:   "HOGE",
		Fuga:   "FUGA",
		FooBar: 12,
	}
	c := &mergeSampleStruct{
		Foo:    "foo",
		Hoge:   "HOGE",
		FooBar: 123,
	}

	d := new(mergeSampleStruct)
	mergeStruct(s, c, d)
	if d.Foo != c.Foo {
		t.Errorf("Foo want %v, got %v", c.Foo, d.Foo)
	}
	if d.Bar != s.Bar {
		t.Errorf("Bar want %v, got %v", s.Bar, d.Bar)
	}
	if d.Hoge != c.Hoge {
		t.Errorf("Hoge want %v, got %v", c.Hoge, d.Hoge)
	}
	if d.Fuga != s.Fuga {
		t.Errorf("Bar want %v, got %v", s.Fuga, d.Fuga)
	}
	if d.FooBar != c.FooBar {
		t.Errorf("FooBar want %v, got %v", c.FooBar, d.FooBar)
	}
}
