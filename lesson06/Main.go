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
	var students [3]string;
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Jisoo"
	students[2] = "Rose"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #2: %v\n", students[1])

	// Determine length
	fmt.Printf("Number of Students: %v\n", len(students))

	// Arrays can be make up of any type but all elements have to be the same type
	var identityMatrix [3][3]int = [3][3]int{ [3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Printf("Identity Matrix: %v\n", identityMatrix )

	// Same thing
	var identityMatrix2 [3][3]int
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[1] = [3]int{0, 1, 0}
	identityMatrix2[2] = [3]int{0, 0, 1}
	fmt.Printf("Identity Matrix2: %v\n", identityMatrix2 )

	// Arrays are considered values (not references)
	a := [...]int{1,2,3}
	b := a
	b[1] = 5
	fmt.Println(a) // [1, 2, 3]
	fmt.Println(b) // [1, 5, 3]

}