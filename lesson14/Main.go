package main

import (
	"fmt"
	"time"
	"sync"
	"runtime"
)

// Goroutines
// - Creating goroutines
// - Synchronization
//   - WaitGroups
//   - Mutexes
// - Parallelism

// - Best Practices
//   - Don't create goroutines in libraries
//     - Let consumer control concurrency
//   - When creating a goroutine, know how it will end
//     - Avoid subtle memory leaks
//   - Check for race conditions at compie time
//     - `go run --race main.go`

// Summary
// - Creating goroutines
//   - Use `go` keyword in front of function call
//   - When using anonymous functions, pass data as local variables
// - Synchronization
//   - Use sync.WaitGroup to wait for groups of goroutines to complete
//   - Use sync.Mutex and sync.RWMutex to protect data access
// - Parallelism
//   - By default, Go will use CPU threads equal to avaiable cores
//   - Change with runtime.GOMAXPROCS
//   - More threads can increase performance, but too many can slow it down
// - Best practices
//   - Don't create goroutines in libraries
//     - Let consume control concurrency
//   - When creating a goroutine, know how it will end
//     - Avoid suble memory leaks
//   - Check for race conditions at compile time

var wg = sync.WaitGroup{}

func main() {
	// example01()
	// example02()
	// example03()
	// example04()
	// example05()
	// example06()
	// example07()
	// example08()

	// GOMAXPROCS can be tuned with values above number of operating system threads for best performance
	//fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))  // interrogate how many thread available
	//runtime.GOMAXPROCS(1) // Set to be single-threaded
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))  // interrogate how many thread available
  example03()
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

// example05 WaitGroup
func example05() {
	var msg = "Hello in example05"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Goodbye"
	wg.Wait()
}

// example06 No consistent ordering. Race condition
var counter = 0

func sayHello06() {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
}

func increment06() {
  	counter++
		wg.Done()
}

func example06() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello06()
		go increment06()
	}
	wg.Wait()
}

// example07 Mutex attempted fix.  Not quite right.
var mutex07 = sync.RWMutex{}
var wg07 = sync.WaitGroup{}

var counter07 = 0

func sayHello07() {
	mutex07.RLock()
	fmt.Printf("Hello #%v\n", counter07)
	mutex07.RUnlock()
	wg07.Done()
}

func increment07() {
	  mutex07.Lock()
  	counter07++
		mutex07.Unlock()
		wg07.Done()
}

func example07() {
	for i := 0; i < 10; i++ {
		wg07.Add(2)
		go sayHello07()
		go increment07()
	}
	wg07.Wait()
}

// example08 Lock mutex out of go routine in a single context
var mutex08 = sync.RWMutex{}
var wg08 = sync.WaitGroup{}

var counter08 = 0

func sayHello08() {
	fmt.Printf("Hello #%v\n", counter08)
	mutex08.RUnlock()
	wg08.Done()
}

func increment08() {
  	counter08++
		mutex08.Unlock()
		wg08.Done()
}

func example08() {
	for i := 0; i < 10; i++ {
		wg08.Add(2)
		mutex08.RLock()
		go sayHello08()
	  mutex08.Lock()
		go increment08()

	}
	wg08.Wait()
}