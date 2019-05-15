package main

import "fmt"

// Package level scope
var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// The compiler assigned values to the variables. What are these values called?
	// Those values are called the Zero Value.
}