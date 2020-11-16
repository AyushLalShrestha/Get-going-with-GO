package main

import (
	"fmt"
	"reflect"
)

func main() {

	tst := "string"
	tst2 := 10
	tst3 := 1.2

	fmt.Println(reflect.TypeOf(tst))
	if reflect.TypeOf(tst).String() == "string" {
		goto label
	}

	fmt.Println(reflect.TypeOf(tst2))
label:
	fmt.Println(reflect.TypeOf(tst3))

}
