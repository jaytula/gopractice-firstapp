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
}