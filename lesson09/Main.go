package main

import (
	"fmt"
)

func main() {
	fmt.Println("Looping")
	// - For statements
	//   - simple loops
	//   - exiting early
	//   - looping through collections

	// Most basic for-loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Double initialization
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	// counter can be manipulated.  best to avoid it.
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}
	}

	// Can omit first initializer in for be declaring it outside
	j := 0
	for ; j < 5; j++ {
		fmt.Println(j)
	}

	// Can also eliminat incrementer
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// For loop dropping the semi-colons with only condition
	k := 0
	for k < 5 {
		fmt.Println(j)
		k++
	}

	// For-loop with nothing. Acts as a while loop. Use 'break' to get out
	k2 := 0
	for {
		fmt.Println(k2)
		k2++
		if k2 == 5 {
			break
		}
	}

	// Using continue with for loops
	for k3 := 0; k3 < 10; k3++ {
		if k3%2 == 0 {
			continue
		}
		fmt.Println(k3)
	}

	// Nested loop
	fmt.Println("Nested Loop")
	var c int
	var d int
	for c = 1; c <= 3; c++ {
		for d = 1; d <= 3; d++ {
			fmt.Println(c * d)
		}
	}

	fmt.Println("Nested Loop. Break out if result is greater or equal to 3")
	out:
	for c = 1; c <= 3; c++ {
		for d = 1; d <= 3; d++ {
			fmt.Println(c * d)
			if c * d >= 3 {
				break out
			}
		}
	}

	fmt.Println("Loop through items of slice using range")
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	fmt.Println("Loop through items of array using range")
	arr := [3]int{1, 2, 3}
	for k, v := range arr {
		fmt.Println(k, v)
	}

	fmt.Println("Loop through items of map using range")
	map1 := map[string]int{
		"California": 39250017,
		"Texas": 27862596,
		"Florida": 20612439,
		"New York": 19745289,
	}
	for k, v := range map1{
		fmt.Println(k, v)
	}

	fmt.Println("Loop through string using range")
	s1 := "hello Go"
	for k, v := range s1 {
		fmt.Println(k, v, string(v))
	}

	// We can also loop through channels. Discussed later.

	fmt.Println("If values are not needed")
	for k := range s1 {
		fmt.Println(k)
	}

	fmt.Println("If keys are not needed")
	for _, v := range s1 {
		fmt.Println(v)
	}


}
