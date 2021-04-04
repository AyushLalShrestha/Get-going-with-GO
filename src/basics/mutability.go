package main

import (
	"fmt"
	"math/rand"
)

type Person struct {
	name string
	age  int
}

func addToChannel(chInts chan int) {
	chInts <- rand.Intn(1000)
}
func addToMap(ageMap map[string]int) {
	ageMap["ayush"] = 24
}

func addToSlice(slc []int) {
	slc = append(slc, 123)
}

func incrementAge(p Person) {
	p.age += 1
	fmt.Printf("Person's age in-method: %d\n", p.age)
}

func main() {
	// Only channels & maps are mutable, nothing else is!!
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

	// Map mutability
	people := make(map[string]int)
	addToMap(people)
	fmt.Println(people)

	// Struct mutability
	person := Person{name: "ayush", age: 24}
	incrementAge(person)
	fmt.Println(person.age)

}
