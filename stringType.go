package main

import "fmt"

type ryanInt int


func main() {
	s := "Hello World!"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", bs) // Prints out the UTF-8 Code Point
	}

	for i, v := range s {
		fmt.Printf("At index position %d we have hex %#x\n", i, v)
		// At index position 0 we have hex 0x48 ...
	}

}