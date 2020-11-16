package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main2() {
	chInts := make(chan int, 2) // OK
	// chInts := make(chan int) Not OK because it would be a blocking call
	chInts <- 123 // This would be Non-blocking for buffered channels
	time.Sleep(10 * time.Second)
	n := <-chInts
	fmt.Printf("n: %d\n", n)
}

// Send can be done before any waits but from another thread
func genInts(chInts chan int) {
	fmt.Println("Writing to the channel")
	chInts <- rand.Intn(1000)
	fmt.Println("Done writing to the channel")
}

func main() {
	chInts := make(chan int)
	for i := 0; i < 2; i++ {
		go genInts(chInts)
	}
	time.Sleep(5 * time.Second)
	n := <-chInts
	fmt.Printf("n: %d\n", n)

	time.Sleep(5 * time.Second)
	n = <-chInts
	fmt.Printf("n: %d\n", n)
}
