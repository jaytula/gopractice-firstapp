package main

import "fmt"

// Package level variable declarations
var m float32 = 68

func main() {
	// Declaring a variable then assigning
	// Use when you assign a variable in a loop
	var i int
	i = 42
	_ = i

	// Declaring a variable and assigning
	// Use when type inference is not precise enough
	var j float32 = 27
	_ = j
	
	// Declaring a variable and assigning with type inference
	// Use whenever possible because it is most compact
	k := 99.
	_ = k
	fmt.Printf("%v, %T\n", j, j);
	fmt.Printf("%v, %T\n", k, k);

	// Print out package level variable
	fmt.Printf("%v, %T\n", m, m)
}