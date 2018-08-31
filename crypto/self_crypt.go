package main

import (
	"unicode"
)

type Cipher interface {
	Encryption(string) string
	Decryption(string) string
}

type ciper []int

func (c ciper) cipherAlgorithm(letters string, shift func(int, int) int) string {
	var shiftedText string
	for _, letter := range letters {
		if !unicode.IsLetter(letter) {
			continue
		}
		shiftDist := c[len(shiftedText)%len(c)]
		s := shift(int(unicode.ToLower(letter)), shiftDist)
		switch {
		case s < 'a':
			s += 'z' - 'a' + 1
		case 'z' < s:
			s -= 'z' - 'a' + 1
		}
		shiftedText += string(s)
	}
	return shiftedText
}

func (c *ciper) Encryption(plainText string) string {
	return c.cipherAlgorithm(plainText, func(a, b int) int { return a + b })
}

func (c *ciper) Decryption(cipherText string) string {
	return c.cipherAlgorithm(cipherText, func(a, b int) int { return a - b })
}

func NewCaesar(key int) Cipher {
	return NewShift(key)
}

func NewShift(shift int) Cipher {
	if shift < -25 || shift > 25 || shift == 0 {
		return nil
	}
	c := ciper([]int{shift})
	return &c
}
