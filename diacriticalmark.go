package main

import (
	"encoding/hex"
	"fmt"
)

func DaiacticalmarkCheck() {
	str := "pͪ"
	for _, s := range str {
		fmt.Print(s)
	}
	fmt.Println()
	h := hex.EncodeToString([]byte(str))
	fmt.Println(h)

	u, _ := hex.DecodeString("cdaa")
	fmt.Println(string(u))

	u, _ = hex.DecodeString("3041")
	fmt.Println(string(u))
}
