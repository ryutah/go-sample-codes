package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func main() {
	yml, err := ioutil.ReadFile("sample.yml")
	if err != nil {
		panic(err)
	}

	payload := make(map[string]interface{})
	if err := yaml.Unmarshal(yml, &payload); err != nil {
		panic(err)
	}
	fmt.Println(payload)
}
