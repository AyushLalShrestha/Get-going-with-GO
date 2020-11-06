package main

import (
	"fmt"
)

func main() {

	a := make(map[string]int)

	a["one"] = 1

	fmt.Println(a)
	fmt.Println(a["name"])
}
