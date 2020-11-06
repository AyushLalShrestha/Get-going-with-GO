package basics

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func bufferedForLoops() {
	buf := bytes.NewBufferString("one\ntwo\nthree\nfour\n")

	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {

				fmt.Print(line)
				break
			}
			fmt.Println(err)
			break
		}
		fmt.Print(line)
	}
}

func channelLoops() {
	queue := make(chan string, 2)

	go func() {
		for elem := range queue {
			fmt.Println(elem)
		}
	}()

	queue <- "one"
	queue <- "two"
	time.Sleep(time.Second * 5)
	close(queue)
}
