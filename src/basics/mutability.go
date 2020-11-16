package main

import (
	"fmt"
	"math/rand"
)

func addToChannel(chInts chan int) {
	chInts <- rand.Intn(1000)
}

func addToSlice(slc []int) {
	slc = append(slc, 123)
}

func main() {
	// Channels mutability
	chInts := make(chan int)
	for i := 0; i < 2; i++ {
		go addToChannel(chInts)
	}
	n := <-chInts
	fmt.Printf("n: %d\n", n)

	select {
	case n := <-chInts:
		fmt.Printf("n: %d\n", n)
	}

	// Slice mutability
	slc := make([]int, 5)
	slices := fmt.Sprint("Slices before func call:", slc)
	addToSlice(slc)
	slices = fmt.Sprint(slices, ", after func call:", slc)
	fmt.Println(slices)

}
