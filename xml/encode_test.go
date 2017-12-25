package main

import (
	"encoding/xml"
	"testing"
)

type foo struct {
	Bars bars
}

type bars struct {
	Bar []bar
}

type bar struct {
	A string
	B string
}

func TestMarshalFromSlice(t *testing.T) {
	b := []bar{
		bar{A: "A1", B: "B1"},
		bar{A: "A2", B: "B2"},
	}
	f := foo{Bars: bars{Bar: b}}
	x, err := xml.MarshalIndent(f.Bars, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("\n%s", string(x))
}

func TestUnMarshalToSlice(t *testing.T) {
	x := `
		<bars>
		  <Bar>
			<A>A1</A>
			<B>B1</B>
		  </Bar>
		  <Bar>
			<A>A2</A>
			<B>B2</B>
		  </Bar>
		</bars>
	`
	f := new(foo)
	if err := xml.Unmarshal([]byte(x), &f.Bars); err != nil {
		t.Fatal(err)
	}
	t.Logf("Foo : %v", f)
}
