package main

import (
	"fmt"
	"regexp"
)

func Regexp() {
	str1 := "Query123_Extract"
	str2 := "Query123_Excel"
	str3 := "Query123"

	rxp := regexp.MustCompile(`^(Query\d+)(?:_.+)?$`)

	sub1 := rxp.FindStringSubmatch(str1)
	sub2 := rxp.FindStringSubmatch(str2)
	sub3 := rxp.FindStringSubmatch(str3)

	fmt.Println(sub1)
	fmt.Println(sub2)
	fmt.Println(sub3)
}
