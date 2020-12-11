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

	// Zero value for all numeric types is '0' or the equivalent of '0' for that type
	intvar := 42
	fmt.Printf("%v, %T\n", intvar, intvar)

	// signed integer types
	// - int8 (-128, 127)
	// - int16 (-32768, 32767)
	// - int32 (-2147483648, 2147483647)
	// - int64

	// unsigned integers example
	var uintvar uint16 = 42
	fmt.Printf("%v, %T\n", uintvar, uintvar)

	// unsigned integer types
	// - uint8 (0, 255)
	// - uint16 (0, 65535)
	// - uint32 (0, 4294967295)
	// - NO uint64 

	// There's also a type `byte` which an alias for `int8`

}