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
	greeting := "Helloo"
	name := "Stacey"
	sayGreeting(greeting, name)
	fmt.Println(name)
	sayGreetingPointers(&greeting, &name)
	fmt.Println(name)
  sum("The sum is ", 1,2,3,4,5)
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

func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Ted"
}

func sayGreetingPointers(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
}

// Variadic parameter. Can only hav one. Has to be at the end. The variadic parameter becomes a slice.
func sum(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}