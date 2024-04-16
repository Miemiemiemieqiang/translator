package main

import (
	"fmt"
	"io"
	"os"
	"translator/mgt"
	"translator/translate"
)

func main() {
	f, err := os.Open("test/json3.json")
	defer f.Close()

	if err != nil {
		panic(err)
	}
	input, err := io.ReadAll(f)

	if err != nil {
		panic(err)
	}
	manager := mgt.DefaultManager
	output, err := manager.Translate(input, translate.JSON, translate.YAML)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
