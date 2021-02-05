package main

import (
	"fmt"
)

// - Basics
// - Composing interfaces
// - Type conversion
//   - The empty interfaces
//   - Type switches
// - Implementing with values vs pointers
// - Best practices

func main() {
	fmt.Println("Interfaces!!!")

	var w Writer = ConsoleWriter{}
	n, err := w.Write([]byte("Hello Go!"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	var myInt =  IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
	  fmt.Println(inc.Increment())
	}
}

// Writer Interfaces describe behavior. So instead of data, we have methods
type Writer interface {
	Write([]byte) (int, error)
}

// ConsoleWriter blah
type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// - Interface naming convention. If the interface has a single method `Read` then the interface name should be `Reader`
// - For interface with multiple methods, you should name it by what it does.
// - Most common way to use interface is to implement them on a struct but it doesn't have to be a struct

// Incrementer increments stuff
type Incrementer interface {
	Increment() int
}

// IntCounter type alias for int
type IntCounter int

// Increment is method on *IntCounter type
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}