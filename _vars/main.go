package main

import "fmt"

var globalName = "Global Name"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Greg"
	var age int = 26
	const isCool = true

	name := "Greg"
	bro1, bro2 := "Jon", "Matt"
	size := 5.11
	fmt.Println(name, age)

	fmt.Println("Height of ", size)
	// Getting types
	fmt.Printf("%T\n", name)

	fmt.Println(bro1, bro2)
}
