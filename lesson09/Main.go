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
}
