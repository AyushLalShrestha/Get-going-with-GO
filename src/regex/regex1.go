package main

import (
	"fmt"
	"regexp"
)

func main() {
	reObj := regexp.MustCompile(`name=(.*),caste=(.*)`)
	re := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.FindAll([]byte(`seafood fool`), -1))
	fmt.Printf("%q\n", reObj.FindAll([]byte(`name=ayush,caste=shrestha`), -1))

}
