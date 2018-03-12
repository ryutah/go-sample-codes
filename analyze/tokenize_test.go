package main

import (
	"strings"
	"testing"

	"github.com/ikawaha/kagome.ipadic/tokenizer"
)

func TestTokenize(t *testing.T) {
	dic := `
明々後日,明々後日,シアサッテ,カスタム名詞
言い間違い,言い間違い,イイマチガイ,カスタム名詞
`

	rDic := strings.NewReader(dic)
	uDicRecord, err := tokenizer.NewUserDicRecords(rDic)
	if err != nil {
		t.Fatal(err)
	}
	uDic, err := uDicRecord.NewUserDic()
	if err != nil {
		t.Fatal(err)
	}

	s := "明々後日はいい天気だといいですね。もし晴れたら、明日は走ろうと思ってます。"

	tknizer := tokenizer.New()
	tknizer.SetUserDic(uDic)
	tokens := tknizer.Tokenize(s)
	for _, token := range tokens {
		t.Log(token.Class)
		t.Log(token.Features())
	}
}
