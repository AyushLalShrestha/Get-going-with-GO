package main

import (
	"fmt"
)

func getMethod(methodName string) func() int {
	var a int
	return func() int {
		a = a + 1
		return a
	}
}

func main() {
	method1 := getMethod("increment")
	fmt.Println(method1())
	fmt.Println(method1())
	variadicFunction("John", "J", "Johny", "Jimothy")
}

func variadicFunction(name string, nickNames ...string) {
	fmt.Println("No. of nick names: ", len(nickNames))
}
