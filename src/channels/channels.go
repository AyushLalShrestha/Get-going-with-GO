package main

import (
	"fmt"
	"time"
)

func main() {

	queue := make(chan string, 2)
	// un-bufferred channel.
	// someone should already be listening before writing on it
	done := make(chan bool)

	go func() {
		// It will wait until the channel is closed
		for elem := range queue {
			fmt.Println(elem, "- waiting for next....")
		}
		fmt.Println("Done, since channel closed")
		done <- true
	}()

	queue <- "one"
	time.Sleep(time.Second * 3)
	queue <- "two"

	time.Sleep(time.Second * 2)
	fmt.Println("Closing the channel now")
	close(queue)
	<-done

}
