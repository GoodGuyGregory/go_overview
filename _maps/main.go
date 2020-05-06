package main

import "fmt"

func main() {
	fmt.Println("Maps are Key value Pairs")
	// makes a map for the key and values allm of type string
	// emails := make(map[string]string)

	// Assign key value
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Tyler"] = "tyler@gmail.com"

	emails := map[string]string{"Tyler": "tyler@gmail.com", "Gretchen": "gretchen@gmail.com"}

	emails["Jake"] = "jake@gmail.com"
	// Display elements in a Map
	fmt.Println(emails)

	// return the length of the map
	fmt.Println(len(emails))

	// return the value from a key
	fmt.Println(emails["Bob"])

	// Delete elements in a map
	delete(emails, "Gretchen")
	fmt.Println("Deleted Bob")
	fmt.Println(emails)
}
