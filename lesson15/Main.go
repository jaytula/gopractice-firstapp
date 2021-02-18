package main

import (
	"fmt"
	"sync"
)


// Agenda
// - Channel basics
// - Restricting data flow
// - Buffered channels
// - Closiing channels
// - For...range loops with channels
// - Select statements

var wg = sync.WaitGroup{}

func main() {
  example01()
}

// Create an int channel and send 42 into it in one goroutine
// Other goroutine receives the int and prins it out.
func example01() {
	ch := make(chan int)
	wg.Add(2)

	// Receiving go routine
	go func() {
		i := <- ch
		fmt.Println(i)
		wg.Done()
	}()

	// Sending go routine
	go func() {
		i := 42
		ch <- i  // passed by value
		i = 27
		wg.Done()
	}()
	wg.Wait()
}