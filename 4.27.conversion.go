/*
	Go is a static programming language, all about type.
	Once you declare a variable the type of the variable will not change.
*/

package main

import "fmt"

var a int

// Here we are creating our own type
type hotdog int
var b hotdog

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// At this point we are assigning a variable to our custom type, hotdog.
	b = 43
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// a = b // ERROR: cannot use b (type hotdog) as type int in assignment
	// We can use conversion in order for this to work. Unlike other languages we do not call this casting.

	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	/*
		Returns:
				43
				int
	*/
}