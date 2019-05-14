package main

import "fmt"

// Declare + Assign = Initialization
var y = 100

// Declares there is a variable with the identifier z of type int, and declares it without a value.
var z int

func lkj() {
	x := 42
	fmt.Println(x)
	x = 99 // Updates value
	fmt.Println(x)
	fmt.Println(y)
	foo()
	fmt.Println(z)
}

func foo() {
	fmt.Println("Again:", y)
}

// Use var and when you need variables inside of function body use the short declaration operator.
// Best practice = Limit the scope of variables and as much as possible use a short declaration operator.
