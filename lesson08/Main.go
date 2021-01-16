package main

import (
	"fmt"
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

	// There is also boolean operators || and &&
	guess = -1
	if guess < 1 || guess > 100 {
		fmt.Println("Number is not between 1 and 100")
	}
}