package main

import "fmt"

var y = 42

func main() {
	// Prints to std.out
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	y = 911
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	// fmt.Sprintf() prints to a string which can then be assigned to a variable.
	s := fmt.Sprintf("%#X\t%b\t%x\n", y, y, y)
	fmt.Println(s)

	// %v gives us the value in the default format.
	fmt.Printf("%v", y)
}

/*
	Logs out:
			42
			int
			101010
			2a
			0x2a
			0x38f
			0x38f   1110001111      38f
			0X38F   1110001111      38f

			911
	There are many different functions within the fmt package which allow us to print data in different ways.
*/

