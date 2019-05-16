package main

import "fmt"

type ryanInt int

const (
	a = 42
	b = 42.78
	c = "james bond"
// or
// const a = 42
// const b = 42.78
// const c = "james bond"


func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}