package main

import "fmt"

func main() {
	fmt.Println("Ranges are used to iterate through arrays ")
	ids := []int{33, 76, 54, 23, 11, 2}

	// Loop through the ids using range
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	fmt.Println("Without the Index")
	// if you do not want to use the variable
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add all the IDs together
	sum := 0
	for _, id := range ids {
		sum += id

	}
	fmt.Println("Sum", sum)

	// range with Map
	emails := map[string]string{"Bernie": "bernie@gmail.com", "Trump": "trump@gmail.com"}
	// key value pairs need to be considered
	for k, v := range emails {
		// %s adds slice elements
		fmt.Printf("%s: %s\n", k, v)
	}

	// Just wanting the keys
	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
