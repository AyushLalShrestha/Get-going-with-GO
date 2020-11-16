package main

import (
	"fmt"
)

func unicodeCodePoint() {
	s := "㈲"
	fmt.Println(s, "=", []byte(s))
	for i := 12850; i < 12852; i++ {
		stringi := string(i)
		bytesStringi := []byte(stringi)
		fmt.Println("Unicode-codepoint=", i, ", character=", stringi, ",", len(bytesStringi), "bytes=", bytesStringi)
	}
}

func main() {
	unicodeCodePoint()
}
