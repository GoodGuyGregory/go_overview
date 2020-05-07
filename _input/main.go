package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//stores space separated values into successive arguments
	// var storageVariable variableType
	// fmt.Scan(&storageVariable) //assuming fmt is imported

	//reads line all in one go
	// reader := bufio.NewReader(os.Stdin) //create new reader, assuming bufio imported
	// var storageString string
	// storageString, _ := reader.ReadString('\n')

	// Scanner functions are used for splitting space-delimited tokens
	var age int
	fmt.Println("What is your age?")
	_, err := fmt.Scan(&age)
	if err == nil {

		// Reader is used to read full lines
		reader := bufio.NewReader(os.Stdin)
		var name string
		fmt.Println("What is your name?")
		// reads until the end of the line
		name, _ = reader.ReadString('\n')

		fmt.Println("Your name is ", name, " and you're ", age)
	} else {
		fmt.Println("Enter a Number", err)
	}
}
