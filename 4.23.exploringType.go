package main

import "fmt"

var x = 42
var w = "James said, Shaken not stirred."
var s = `
	Shaken
	not
	stirred.
`

//  Strings can be declared by using double quotes and backticks.

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	// 42
	// int
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	// Shaken not stirred.
	// string
}

//  Primitive data types = Basic / Built-In types.
//  Composite data types = Compose data together. Sometimes called a structure or aggregate data types.
