package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const filepath string = "/Users/ayushshrestha/my_projects/GetGoing/src/fileIO/dummyfile.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Open the file pointer
	f, err := os.Open(filepath)
	defer f.Close()
	check(err)

	// Read 1st 5 bytes of the file
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Set the pointer 6 places from the current location & read 2 bytes
	o2, err := f.Seek(6, 1)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Println(string(b2)) // fmt.Printf("%v\n", string(b2[:n2]))

	// Get the current position of the pointer
	offset, err := f.Seek(0, io.SeekCurrent)
	fmt.Println("Curret position of the pointer = ", offset)

	// o3, err := f.Seek(6, 0)
	// check(err)
	// b3 := make([]byte, 2)
	// n3, err := io.ReadAtLeast(f, b3, 2)
	// check(err)
	// fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// _, err = f.Seek(0, 0)
	// check(err)

	// r4 := bufio.NewReader(f)
	// b4, err := r4.Peek(5)
	// check(err)
	// fmt.Printf("5 bytes: %s\n", string(b4))
}

func readAllTheFile() {
	dat, err := ioutil.ReadFile(filepath)
	check(err)
	fmt.Print(string(dat))
}
