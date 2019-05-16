package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)
const (
	d = iota
	e = iota
	f = iota
)

func main() {
	fmt.Println(a) // 0
	fmt.Println(b) // 1
	fmt.Println(c) // 2
	fmt.Printf("%T\n", a) // int
	fmt.Printf("%T\n", b) // int
	fmt.Printf("%T\n", c) // int

	fmt.Println(d) // 0
	fmt.Println(e) // 1
	fmt.Println(f) // 2
	fmt.Printf("%T\n", d) // int
	fmt.Printf("%T\n", e) // int
	fmt.Printf("%T\n", f) // int
}