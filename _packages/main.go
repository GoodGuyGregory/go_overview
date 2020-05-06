package main

import (
	"fmt"
	"math"

	// import module created from path to use the function
	"github.com/travery-media-crashcourse/_packages/strutil"
)

func main() {
	fmt.Println("Packages")
	fmt.Println("Rounding with Math Package")
	fmt.Println(math.Floor(3.8))
	fmt.Println(math.Ceil(6.8))
	fmt.Println(math.Sqrt(144))
	// use reverse function
	fmt.Println(strutil.Reverse("olleh"))

}
