package main

import "fmt"

type Cat struct {
	color string
	breed string
}

type X interface {
	foo()
	bar()
}

func (c Cat) foo() {
	fmt.Println(c.color, c.breed)
}

func (c *Cat) bar() {
	fmt.Println(c.breed, c.color)
}

func main() {
	c := Cat{color: "black", breed: "thoroughbred"}
	var p *Cat
	p = &c

	var x X

	x = p // 0K: *Cat has explicit method bar() and implicit method foo()
	x.foo()

	// x = c // Not OK: c hasn't implemented X
}

// implicit pointer methods satisfy interfaces,
//  but implicit non-pointer methods do not
