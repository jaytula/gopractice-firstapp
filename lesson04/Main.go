package main

import (
	"fmt"
)

/*
 - Boolean type
 - Numeric types
	 - Integers
	 - Floating point
	 - Complex numbers
 - Text types
 */

func main() {
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	a := 1 == 1
	b := 1 == 2
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)

	// This is okay. bool-type has a zero value of false
	var c bool
	fmt.Printf("%v, %T\n", c, c)


}