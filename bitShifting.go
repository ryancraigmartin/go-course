package main

import "fmt"

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)
	y := x << 1 // Shifts the bit over.
	fmt.Printf("%d\t\t%b\n", y, y)

	//4               100
	//8               1000
}