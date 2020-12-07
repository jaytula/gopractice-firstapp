package main

import (
	"fmt"
)

var i int = 27

func main() {
	// This prints out the package-level variable
	fmt.Println(i)

	// The following variable declaration shadows the package-level variable
	var i int = 42
	fmt.Println(i)
}