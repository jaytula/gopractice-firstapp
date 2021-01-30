package main

import (
	"fmt"
)

// Basic syntax
// Parameters
// Return values
// Anonymous functions
// Functions as types
// Methods

// Entrypoint is always in 'package main'
// - Function start with 'func' keyword
// - Capitalized function names means the function is visible outside the package
func main() {
	fmt.Println("Hello, functions")
	sayMessage("Hello Go!")
	sayMessageTwoParams("Hello Two Params", 27)
}

// Example of a function with one parameter
func sayMessage(msg string) {
	fmt.Println(msg)
}

// Function with two parameters
func sayMessageTwoParams(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}