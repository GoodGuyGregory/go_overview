package main

import "fmt"

func main() {
	fmt.Println("Arrays")
	// Arrays must be a fixed length, which can be a problem

	// CREATE AND APPEND
	// var fruitArr [2]string

	// // assign values to fruit array
	// fruitArr[0] = "strawberry"
	// fruitArr[1] = "Bananna"

	// CREATE AND INITIALIZE
	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)
	fmt.Println("Gotta have my " + fruitArr[1] + " with my coffee")

	fruitSlice := []string{"GrapeFruit", "Kiwi", "Grapes"}

	fmt.Println(fruitSlice)

	// return the number of elements
	fmt.Println(len(fruitSlice))
	// Returns a range in the slice
	// Starts at 1 and ends at 2
	fmt.Println(fruitSlice[1:2])

}
