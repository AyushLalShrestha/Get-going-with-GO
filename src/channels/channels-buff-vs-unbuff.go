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

// "Send" can be done before any waits from another thread
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

/*
In Buffered channels (make(chan int, 2)), value can be sent without blocking,
until it is full!
In unbuffered channels (make(chan int)), deadlock can be triggerred while
sending a value to the channel if there is no chance that someone will
read from the channel


An unbuffered channel means that read and writes from and to the channel are blocking.
In a select statement:
1. the read would work if some other goroutine was currently blocked in writing to the channel
2. the write would work if some other goroutine was currently blocked in reading to the channel
3. otherwise the default case is executed, which happens in your case A.
*/
