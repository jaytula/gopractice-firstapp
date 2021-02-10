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
  go sayHello() // Prefix with 'go' to spin off into a green thread. Light-weight abstraction for a thread
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}