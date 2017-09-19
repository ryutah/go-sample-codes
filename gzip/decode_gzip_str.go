package main

import (
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

func main() {
	str := `H4sIAAAJbogA/3SPT0/EIBDFv8ucqy1NVrM9e/G+nnTTsDDbZS1/hEHTNHx3GW4eDAcy837w3ttBeXc1S46SjHcw7bB6qfnWmMi4tj7Jy4ptJ0kmpNcKAMUt23Cb9QU6CNHfUf0RRvFwD4OoIvHzJp1smMdBPA9HcRBPBzEe58/8I52H0sHV4KpfcDXWEMZKd/V8ZU/VGj6gDknd0EoO0tgE0/sOTloGstJstQUeEkXjFijn+m3yOSp8i4ZxWNLU91xx5i79v4EeVfqGcynlFwAA//8=`

	encoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}
	decoded, err := decompressFromGzip(string(encoded))
	if err != nil {
		panic(err)
	}
	fmt.Println(decoded)
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
