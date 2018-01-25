package main

import (
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/lestrrat/go-ngram"
)

func createNgram(str string, n int) []string {
	var (
		newstr  = str
		ret     = make([]string, 0, len(str))
		size    = 0
		runeidx = make([]int, 1, len(str))
	)

	for len(newstr) > 0 {
		_, wide := utf8.DecodeRuneInString(newstr)
		size += wide
		runeidx = append(runeidx, size)
		newstr = newstr[wide:]
	}

	for i, j := 0, n; j < len(runeidx); j++ {
		left, right := runeidx[i], runeidx[j]
		s := str[left:right]
		ret = append(ret, s)
		i = j - (n - 1)
	}
	return ret
}

func createNgramByLib(s string, n int) []string {
	tokens := ngram.NewTokenize(n, s).Tokens()
	ret := make([]string, 0, len(tokens))
	for _, token := range tokens {
		ret = append(ret, token.String())
	}
	return ret
}

func createNgramByRune(s string, n int) []string {
	runes := []rune(s)
	grams := make([]string, 0, len(runes))
	for left, right := 0, n; right <= len(runes); right++ {
		grams = append(grams, string(runes[left:right]))
		left++
	}
	return grams
}

func BenchmarkCreateNgram(b *testing.B) {
	str := strings.Repeat("あ", 1000)
	b.Run("Myself", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			createNgram(str, 2)
		}
	})
	b.Run("Lib", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			createNgramByLib(str, 2)
		}
	})
	b.Run("Rune", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			createNgramByRune(str, 2)
		}
	})
}

func TestCreateNgram_Bigram_Multibyte(t *testing.T) {
	str := "ハローワールド"
	want := []string{
		"ハロ",
		"ロー",
		"ーワ",
		"ワー",
		"ール",
		"ルド",
	}
	got := createNgram(str, 2)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("\nwant : %v\n got : %v", want, got)
	}
}

func TestCreateNgram_Bigram(t *testing.T) {
	str := "Hello, World"
	want := []string{
		"He",
		"el",
		"ll",
		"lo",
		"o,",
		", ",
		" W",
		"Wo",
		"or",
		"rl",
		"ld",
	}
	got := createNgram(str, 2)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("\nwant : %v\n got : %v", want, got)
	}
}

func TestCreateNgram_Unigram_Multibyte(t *testing.T) {
	str := "ハローワールド"
	want := []string{
		"ハ",
		"ロ",
		"ー",
		"ワ",
		"ー",
		"ル",
		"ド",
	}
	got := createNgram(str, 1)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("\nwant : %v\n got : %v", want, got)
	}
}

func TestCreateNgram_Unigram(t *testing.T) {
	str := "Hello, World"
	want := []string{
		"H",
		"e",
		"l",
		"l",
		"o",
		",",
		" ",
		"W",
		"o",
		"r",
		"l",
		"d",
	}
	got := createNgram(str, 1)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("\nwant : %v\n got : %v", want, got)
	}
}

func TestCreateNgram_Trigram_Multibyte(t *testing.T) {
	str := "ハローワールド"
	want := []string{
		"ハロー",
		"ローワ",
		"ーワー",
		"ワール",
		"ールド",
	}
	got := createNgram(str, 3)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("\nwant : %v\n got : %v", want, got)
	}
}

func TestCreateNgram_Trigram(t *testing.T) {
	str := "Hello, World"
	want := []string{
		"Hel",
		"ell",
		"llo",
		"lo,",
		"o, ",
		", W",
		" Wo",
		"Wor",
		"orl",
		"rld",
	}
	got := createNgram(str, 3)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("\nwant : %v\n got : %v", want, got)
	}
}

func TestCreateNgram_LessGrams(t *testing.T) {
	str := "H"
	want := []string{}
	got := createNgram(str, 3)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("\nwant : %v\n got : %v", want, got)
	}
}

func TestCreateNgramByRune_Bigram_Multibyte(t *testing.T) {
	str := "ハローワールド"
	want := []string{
		"ハロ",
		"ロー",
		"ーワ",
		"ワー",
		"ール",
		"ルド",
	}
	got := createNgramByRune(str, 2)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("\nwant : %v\n got : %v", want, got)
	}
}

func TestCreateNgramByRune_Bigram(t *testing.T) {
	str := "Hello, World"
	want := []string{
		"He",
		"el",
		"ll",
		"lo",
		"o,",
		", ",
		" W",
		"Wo",
		"or",
		"rl",
		"ld",
	}
	got := createNgramByRune(str, 2)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("\nwant : %v\n got : %v", want, got)
	}
}

func TestCreateNgramByRune_Unigram_Multibyte(t *testing.T) {
	str := "ハローワールド"
	want := []string{
		"ハ",
		"ロ",
		"ー",
		"ワ",
		"ー",
		"ル",
		"ド",
	}
	got := createNgramByRune(str, 1)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("\nwant : %v\n got : %v", want, got)
	}
}

func TestCreateNgramByRune_Unigram(t *testing.T) {
	str := "Hello, World"
	want := []string{
		"H",
		"e",
		"l",
		"l",
		"o",
		",",
		" ",
		"W",
		"o",
		"r",
		"l",
		"d",
	}
	got := createNgramByRune(str, 1)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("\nwant : %v\n got : %v", want, got)
	}
}

func TestCreateNgramByRune_Trigram_Multibyte(t *testing.T) {
	str := "ハローワールド"
	want := []string{
		"ハロー",
		"ローワ",
		"ーワー",
		"ワール",
		"ールド",
	}
	got := createNgramByRune(str, 3)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("\nwant : %v\n got : %v", want, got)
	}
}

func TestCreateNgramByRune_Trigram(t *testing.T) {
	str := "Hello, World"
	want := []string{
		"Hel",
		"ell",
		"llo",
		"lo,",
		"o, ",
		", W",
		" Wo",
		"Wor",
		"orl",
		"rld",
	}
	got := createNgramByRune(str, 3)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("\nwant : %v\n got : %v", want, got)
	}
}

func TestCreateNgramByRune_LessGrams(t *testing.T) {
	str := "H"
	want := []string{}
	got := createNgramByRune(str, 3)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("\nwant : %v\n got : %v", want, got)
	}
}
