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
  // example01()
	// example02()
	// example03()
	example05()
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

// example02 deadlock example
func example02() {
	ch := make(chan int)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	for j := 0; j < 5; j++ {
		wg.Add(2)

    go func() {
			ch <- 42  // unbuffered channel. will stop execution until it can be received.  first time yes but second time will deadlock
			wg.Done()
		}()
	}
	wg.Wait()
}

// example03 - modified from example01
func example03() {
	ch := make(chan int)
	wg.Add(2)

	// 2. Receives 42 from channel and prints it out
	// 3. Send 27 into channel
	go func() {
		i := <- ch
		fmt.Println(i)
		ch <- 27
		wg.Done()
	}()

	// 1. Send 42 into the channel
	// 4. Receives 27 from channel and prints it out
	go func() {
		ch <- 42
    fmt.Println(<-ch)
		wg.Done()
	}()
	wg.Wait()
}

// example04
func example04() {
	ch := make(chan int) // bi-direction channel
	wg.Add(2)

	// ch is a receive only channel
	go func(ch <-chan int) {
		i := <- ch
		fmt.Println(i)
		// ch <- 27 would be error
		wg.Done()
	}(ch)


	// ch param is a send-only channel
	go func(ch chan<- int) {
		ch <- 42
    // fmt.Println(<-ch) would be error
		wg.Done()
	}(ch)
	wg.Wait()
}

// example05 - deadlocks
func example05() {
	ch := make(chan int, 50) // Add buffer. Creates a buffer that can store 50 integers. Works around deadlock.
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)  // Added to receive the other integer 27 sent to the channel below
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27  // Buffered channel avoids deadlock but we loose this.
		wg.Done()
	}(ch)
	wg.Wait()
}