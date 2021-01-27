package main

import (
	"fmt"
)

// AGENDA
// - Creating pointers
// - Dereferencing pointers
// - The new function
// - Working with nil
// - Types with internal pointers

func main() {
	a := 42
	b := a
	fmt.Println(a, b) // 42 42
	a = 27
	fmt.Println(a, b) // 27 42

	// Address-of and Dereferencing
	var a2 int = 42
	var b2 *int = &a2
	fmt.Println(a2, b2) // 42 0xc0000180a8
	fmt.Println(&a2, b2) // 0xc0000180a8 0xc0000180a8
	fmt.Println(a2, *b2) // 42 42
	a2 = 27
	fmt.Println(a2, *b2) // 27 27
	*b2 = 14
	fmt.Println(a2, *b2) // 14 14

	a3 := [3]int{1, 2, 3}
	b3 := &a3[0]
	c3 := &a3[1]
	// c3 := &a3[1] + 1 // Pointer arithmtetic not allowed.  Use unsafe package if needed.

	fmt.Printf("%v %p %p\n", a3, b3, c3) // [1 2 3] 0xc000130020 0xc000130028 (diff in address indicatrs 8-bytes apart)
}