package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func foo() {
	for i := 0; i < 25; i++ {
		fmt.Println("foo:", i)
		time.Sleep(30 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 25; i++ {
		fmt.Println("bar:", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("No. of CPUs = ", runtime.NumCPU())
}
