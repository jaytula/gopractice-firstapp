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

}