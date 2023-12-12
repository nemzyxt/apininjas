package main

import (
	"fmt"
	"github.com/nemzyxt/apininjas/dictionary"
)

func main() {
	dict := dictionary.NewClient("y7Hf8C0WWUrrBjcz3wbzpcN77ZUqw4OVsc1c4jDE")
	resp, err := dict.CheckWord("peace")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
