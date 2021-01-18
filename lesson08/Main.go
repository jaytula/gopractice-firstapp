package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Control Flow")

	// Simplest If stateement
	if true {
		fmt.Println("The test is true")
	}

	// Initializer syntax
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas": 27862596,
	}
	if pop, ok := statePopulations["Texas"]; ok {
		fmt.Println(pop)
	}
	if pop, ok := statePopulations["NotOk"]; ok {
		fmt.Println(pop)
	}

	// Less-than, greater-than, equality operators
	number := 50
	guess := 30
	if guess < number {
		fmt.Println("Too low")
	}
	if guess > number {
		fmt.Println("Too high")
	}
	if guess == number {
		fmt.Println("You got it!")
	}
	guess = 50
	fmt.Println(number<=guess, number>=guess, number!=guess)
	guess = 30
	fmt.Println(number<=guess, number>=guess, number!=guess)
	guess = 70
	fmt.Println(number<=guess, number>=guess, number!=guess)

	// There are boolean operators || and && and !
	guess = -1
	if guess < 1 || guess > 100 {
		fmt.Println("Number is not between 1 and 100")
	}

	// Example && usage
	guess = 100
	if guess >= 1 && guess <= 100 {
		fmt.Println("Number is between 1 and 100")
	}

	guess = 96
	// Example usage of NOT ! operator
	if !(guess < 95) {
		fmt.Printf("Number %d is not less than 95\n", guess)
	}

	// Short-circuit example with or
	guess = 1
	if guess == 1 || returnTrue() { // returnTrue is not run because we already got our true
		fmt.Println("guess is 1")
	}
	guess = 0
	if guess == 1 || returnTrue() {
		fmt.Println("guess is 1")
	}

	// If-else example
	guess = 1
	if guess == 2 {
		fmt.Println("Guess is two")
	} else {
		fmt.Println("Guess is not two")
	}

	// Else if
	guess = 30
	if guess < 1 {
		fmt.Println("The guess must be greater than 1")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100")
	} else {
		fmt.Println("Between")
	}

	// Floating point number calculations are inexact
	myNum := 0.123
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
		fmt.Printf("%f %f\n", myNum, math.Pow(math.Sqrt(myNum), 2))
	}

	// Better way
	myNum = 0.123456789
	if math.Abs(myNum / math.Pow(math.Sqrt(myNum), 2) - 1) < 0.0001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
		fmt.Printf("%f %f\n", myNum, math.Pow(math.Sqrt(myNum), 2))
	}

	// Switch statement - simple example
	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}