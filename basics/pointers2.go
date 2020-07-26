package main

import "fmt"

func max(a *int, b *int, p **int) {
	if *a > *b {
		*p = a
	} else {
		*p = b
	}
}

func main() {
	i := 3
	j := 5
	var k *int
	max(&i, &j, &k) // assigns &j to k
	fmt.Println(*k)
}
