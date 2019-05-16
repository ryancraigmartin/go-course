package main

import "fmt"

type ryanInt int

var x ryanInt

func main() {
	fmt.Println(x)
	fmt.Printf("%T", x)
	x = 42
	fmt.Println(x)
}
