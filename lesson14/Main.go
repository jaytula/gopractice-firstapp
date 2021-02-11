package main

import (
	"fmt"
	"time"
)

// Goroutines
// - Creating goroutines
// - Synchronization
//   - WaitGroups
//   - Mutexes
// - Parallelism
// - Best Practices

func main() {
	// example01()
	// example02()
	example03()
	example04()
}

func sayHello() {
	fmt.Println("Hello")
}

func example01() {
	go sayHello() // Prefix with 'go' to spin off into a green thread. Light-weight abstraction for a thread
	time.Sleep(100 * time.Millisecond)
}

// example02 Immediately Invoked Function Expression
func example02() {
	var msg = "Hello IIFE"
	go func() {
		fmt.Println(msg)
	}()
	time.Sleep(100 * time.Millisecond)
}

// example03 This prints out 'Goodbye'
func example03() {
	var msg = "Hello IIFE"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}

// example04 Fix
func example04() {
	var msg = "Hello example04"
	go func(s string) {
		fmt.Println(s)
	}(msg)
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}