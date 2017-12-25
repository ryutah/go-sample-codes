package lib

//go:generate /Users/ryuta/Project/src/github.com/ryutah/mysamples/generate/generate -hoge aaaa a b c

type Foo interface {
	Bar(string) string
}
