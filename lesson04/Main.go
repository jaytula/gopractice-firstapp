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

	// Example of arithmetic
	num1 := 10
	num2 := 3
	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)

	// We cannot add different types. compile time error
	var num3 int = 10
	var num4 int8 = 3
	// fmt.Println(num3 + num4)

	// Type conversion to get it to work
	fmt.Println(num3 + int(num4))
	
	// Bit operators
	num5 := 10 // 1010
	num6 := 3  // 0011

	fmt.Println(num5 & num6) // AND 0010 = 2
	fmt.Println(num5 | num6) // OR 1011 = 11
	fmt.Println(num5 ^ num6) // XOR 1001 = 9
	fmt.Println(num5 &^ num6) // num5 AND-NOT num6 1010 & 1100 = 1000 = 8

	// Bit shifting
	num7 := 8
	fmt.Println(num7 << 3) // 64
	fmt.Println(num7 >> 3) // 1

	// Floating-point types
	// - float32
	// - float64
	myfloat := 3.14  // Initializer syntax is float64
	myfloat = 13.7e72
	myfloat = 2.1E14
	fmt.Printf("%v, %T\n", myfloat, myfloat)
	
	var myfloat2 float32 = 6.28
	fmt.Printf("%v, %T\n", myfloat2, myfloat2)

	// Arithmetic operation on floating point types
	fnum1 := 10.2
	fnum2 := 3.7
	fmt.Println(fnum1 + fnum2)
	fmt.Println(fnum1 - fnum2)
	fmt.Println(fnum1 * fnum2)
	fmt.Println(fnum1 / fnum2)

	// Complex type (Imaginary numbers)
	var cnum1 complex64 = 1 + 2i
	var cnum2 complex64 = 2i
	cnum3 := 2 + 3i

	fmt.Printf("%v, %T\n", cnum1, cnum1)
	fmt.Printf("%v, %T\n", cnum2, cnum2)
	fmt.Printf("%v, %T\n", cnum3, cnum3)	

	// Operations on Complex type
	cnum4 := 1 + 2i
	cnum5 := 2 + 5.2i
	fmt.Println(cnum4 + cnum5)
	fmt.Println(cnum4 - cnum5)
	fmt.Println(cnum4 * cnum5)
	fmt.Println(cnum4 / cnum5)

	// Pulling out with `real` and `imag` builtin functions
	fmt.Printf("%v, %T\n", real(cnum4), real(cnum4))
	fmt.Printf("%v, %T\n", imag(cnum4), imag(cnum4))

	// Making a complex number
	var cnum6 complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", cnum6, cnum6)

}