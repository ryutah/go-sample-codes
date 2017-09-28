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
	str := `H4sIAAAJbogA/1SPO08DMQzHv4vnK/dgQNzMwl4mqE5p4l4DlweJA6pO+e7YEUuVIbL90/+xgw7+YteSFNngYd5hC8rIbzCT9W19VOcN206RykivDAClW3HxupgzdBBT+ER9dzgQCwzj4f/EEIlMQ44uLtMwPg3P07B8lV/lwwS1g4vFzbzgZp0lTAx2/L5LIHaHD+Ah6ys6JVkam2F+38ErJ0DRRlxuUYZMyfoV6ollcyhJ41uygsOa576XlovU6VuOaXx80PkHTrXWPwAAAP//`

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
