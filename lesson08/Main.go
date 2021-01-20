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

	// Instead of fall-through case. We can specify multiple items in a case
	switch 7 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four, or six")
	default:
		fmt.Println("something else")
	}

	// Can use an initializer for switch value
	switch i1:= 2 + 3; i1 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four, or six")
	default:
		fmt.Println("something else")
	}

	// Another syntax for case, 'tag-less'
	j := 11
	switch {
	case j <= 10:
		fmt.Println("Less than or equal 10")
	case j <= 20:
		fmt.Println("less than or equal 20")
	default:
		fmt.Println("Greater than 20")
	}

	// Go switch has implicit break.  Use fallthrough keyword
	// Note fallthrough is logic-less. It will fall-through no matter what.
	j = 10
	switch {
	case j <= 10:
		fmt.Println("Less than or equal 10")
		fallthrough
	case j >= 20:
		fmt.Println("less than or equal 20")
	default:
		fmt.Println("Greater than 20")
	}

	// Example of type-switch
	var iface1 interface{} = "asdf"
	typeSwitch(iface1)
	typeSwitch(2)
	typeSwitch(2.4)
	typeSwitch([3]int{1, 2, 3})
}

func typeSwitch(iface1 interface{}) {
	switch iface1.(type) {
	case int:
		fmt.Println("iface1 is an int")
		// break
		fmt.Println("This will print too")
	case float64:
		fmt.Println("iface1 is a float64")
	case string:
		fmt.Println("iface1 is a string")
	case [3]int:
	  fmt.Println("iface1 is an array of three values")
	default:
		fmt.Println("iface1 is another type")
	}
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}