package main

import (
	"fmt"
	"github.com/Miemiemiemieqiang/translator/mgt"
	"github.com/Miemiemiemieqiang/translator/translate"
	"io"
	"os"
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
	output, err := manager.Translate(input, translate.JSON, translate.CSV)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
