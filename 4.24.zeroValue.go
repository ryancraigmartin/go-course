package main

import "fmt"

var u string
var t int

func main() {
	// Declare a variable to be of a certain type
	// and then assign a value of that type to the variable.
	fmt.Println("Printing the value of u", u, "ending")
	fmt.Printf("%T\n", u)

	u = "Bond, James Bond"

	fmt.Println(u)
	fmt.Printf("%T\n", u)

	fmt.Println(t)
	fmt.Printf("%T\n", z)

}
