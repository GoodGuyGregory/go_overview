package main

import "fmt"

func main() {
	fmt.Println("Loops")
	// Long Method:
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//  Short Method:
	for i := 1; i < 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// FizzBuzz
	for i := 1; i < 100; i++ {
		if i%15 == 0 {
			fmt.Printf("%d FizzBuzz", i)
		} else if i%3 == 0 {
			fmt.Printf("%d Fizz ", i)
		} else if i%5 == 0 {
			fmt.Printf("%d Buzz ", i)
		} else {
			fmt.Println(i)
		}
	}
}
