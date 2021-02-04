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