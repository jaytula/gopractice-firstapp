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

	var ms *myStruct
	fmt.Println(ms) // <nil>
	ms = &myStruct{foo: 42}
	fmt.Println(ms) // &{42}

	ms = new(myStruct)
	fmt.Println(ms) // &{0}
	(*ms).foo = 27
	fmt.Println(ms) // &{27}
	fmt.Println((*ms).foo) // 27
	ms.foo = 42
	fmt.Println(ms.foo)

	// Copy by value
	a4 := [3]int{1, 2, 3}
	b4 := a4  // A copy
	fmt.Println(a4, b4)
	a4[1] = 42
	fmt.Println(a4, b4) // [1 42 3] [1 2 3]

	// Slice by reference
	a5 := []int{1, 2, 3}
	b5 := a5  
	fmt.Println(a5, b5)
	a5[1] = 42
	fmt.Println(a5, b5) // [1 42 3] [1 42 3]

	// Maps are also by reference
	m1 := map[string]string {"foo": "bar", "baz": "buz"}
	m2 := m1
	fmt.Println(m1, m2)
	m1["foo"] = "quz"
	fmt.Println(m1, m2) // Both changed
}

type myStruct struct {
	foo int
}