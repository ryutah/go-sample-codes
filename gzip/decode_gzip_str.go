package main

import (
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func main() {
	str := `H4sIAAAJbogA/1SPPWvEMAyG/4vmpPlYDjIXSvd0ao/DtZ2cSmylllw4gv97rdCleDB634fH8gGW4oJrTkaQIkwHbGSc3s6zYDzj2Xxu/syMGPbyWgGQ9Mhhv9806wdoYE/05e2/spUq6Yf2r6qQqOpE5rDf5pfhAqWBBf3mnv2GAcWnWjb1fGeS+ip8QB3Y3n0wusPJMkzvB0QTFMjWqfmx68CSMK5QrlXLlJP1bwkVh5WnrluJHLcYF+r0o61u3439cOnHsX+y/APXUsovAAAA//8=`

	encoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}
	decoded, err := decompressFromGzip(string(encoded))
	if err != nil {
		panic(err)
	}
	prettyPrintJSON(decoded)
}

func decompressFromGzip(s string) (orig string, err error) {
	strR := strings.NewReader(s)
	r, err := gzip.NewReader(strR)
	if err != nil {
		return
	}

	for {
		p := make([]byte, 4*1024)
		n, e := r.Read(p)
		if e == io.EOF || n <= 0 {
			break
		} else if err = e; err != nil {
			return
		}
		orig += string(p[:n])
	}

	return
}

func prettyPrintJSON(s string) {
	j := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &j); err != nil {
		panic(err)
	}
	pp, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(pp))
}
