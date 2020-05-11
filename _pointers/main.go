package main

import "fmt"

func main() {
	fmt.Println("Pointers in GO")

	a := 5
	// reference to memory address on the memory
	b := &a

	fmt.Println(a, b)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// used remove values from memory address
	fmt.Println(*b)
	fmt.Println(*&a)

	// change value with the pointer
	*b = 10
	// Changes the value at the memory address to 10
	fmt.Println(a)
}
