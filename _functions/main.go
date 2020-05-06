package main

import "fmt"

// Functions in GO
//  func function(<name> <type>) <return type> {

// }
func greeting(name string) string {

	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("Functions")
	fmt.Println(greeting("Greg"))

	fmt.Println("Function that uses ints and returns their some")
	fmt.Println(getSum(22, 64))
}
