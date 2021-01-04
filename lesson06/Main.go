package main

import (
	"fmt"
)

func main() {
	fmt.Println("Lesson 06 - Arrays and Slices")
	/*
		Arrays
			- Creation
			- Built-in functions
			- Working with arrays

		Slices
			- Creation
			- Built-in functions
			- Working with slices
	*/

	// Initialize an array
	grades := [3]int{56, 47, 89}
	fmt.Printf("Grades: %v\n", grades)

	// Initialize an array literal with size implicitly determined
	grades2 := [...]int{56, 47, 89}
	fmt.Printf("Grades: %v\n", grades2)

	// Declare empty array and assign afterwards
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Jisoo"
	students[2] = "Rose"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #2: %v\n", students[1])

	// Determine length
	fmt.Printf("Number of Students: %v\n", len(students))

	// Arrays can be make up of any type but all elements have to be the same type
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Printf("Identity Matrix: %v\n", identityMatrix)

	// Same thing
	var identityMatrix2 [3][3]int
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[1] = [3]int{0, 1, 0}
	identityMatrix2[2] = [3]int{0, 0, 1}
	fmt.Printf("Identity Matrix2: %v\n", identityMatrix2)

	// Arrays are considered values (not references)
	a := [...]int{1, 2, 3}
	b := a // b is assigned a copy of array a
	b[1] = 5
	fmt.Println(a) // [1, 2, 3]
	fmt.Println(b) // [1, 5, 3]

	c := &a
	c[2] = 89
	fmt.Println(a) // [1, 2, 89]
	fmt.Println(c) // &[1, 2, 89]

	// Arrays have a fixed size at compile-time, thus limitting their usefulness

	// Arrays often used to back a slice
	mySliceA := []int{1, 2, 3}
	fmt.Println(mySliceA)
	fmt.Printf("Length: %v\n", len(mySliceA))

	// With slices, we also have a capacity
	fmt.Printf("Capacity: %v\n", cap(mySliceA)) // Lenght of the underlying backing array

	// Slices are refernece types
	mySliceB := mySliceA
	mySliceB[1] = 5 // Changes both
	fmt.Println(mySliceA)
	fmt.Println(mySliceB)

	// Other was to create a slice
	sa := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sb := sa[:]     // slice of all elements
	sc := sa[3:]    // slice from the 4th element to the end
	sd := sa[:6]    // slice first 6 elements
	se := sa[3:6]   // slice of the 4th 5th and 6th elements
	fmt.Println(sa) // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(sb) // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(sc) // [4 5 6 7 8 9 10]
	fmt.Println(sd) // [1 2 3 4 5 6]
	fmt.Println(se) // [4 5 6]

	// Changing changes because slice is like a reference type
	sa[5] = 42
	fmt.Println(sa) // [1 2 3 4 5 42 7 8 9 10]
	fmt.Println(sb) // [1 2 3 4 5 42 7 8 9 10]
	fmt.Println(sc) // [4 5 42 7 8 9 10]
	fmt.Println(sd) // [1 2 3 4 5 42]
	fmt.Println(se) // [4 5 42]

	// - You can slice an array or a slice

	// Using make. 1st argument is the type, 2nd is the length
	ma := make([]int, 3)
	fmt.Println(ma) // [0 0 0]
	fmt.Printf("Length: %v\n", len(ma))
	fmt.Printf("Capacity: %v\n", cap(ma))

	// Make can have a third argumeent for capacity
	mb := make([]int, 3, 100)
	fmt.Println(mb) // [0 0 0]
	fmt.Printf("Length: %v\n", len(mb))
	fmt.Printf("Capacity: %v\n", cap(mb))

	// Slice of int with no elements
	s0 := []int{}
	fmt.Printf("%v\n", s0)
	fmt.Printf("Length %v\n", len(s0))
	fmt.Printf("Capacity %v\n", cap(s0))

	// Use append to add to slice
	s0 = append(s0, 1)
	fmt.Printf("%v\n", s0)
	fmt.Printf("Length %v\n", len(s0))
	fmt.Printf("Capacity %v\n", cap(s0))

	s0 = append(s0, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("%v\n", s0)
	fmt.Printf("Length %v\n", len(s0))
	fmt.Printf("Capacity %v\n", cap(s0))

	// Concatenating slices with spread operator
	s1 := append(s0, []int{1, 2, 3, 4}...)
	fmt.Printf("%v\n", s1)
	fmt.Printf("Length %v\n", len(s1))
	fmt.Printf("Capacity %v\n", cap(s1))

}
