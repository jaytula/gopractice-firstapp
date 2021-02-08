package main

import (
	"fmt"
	"bytes"
	"io"
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

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello YouTube listeners, this is a test"))
	wc.Close()

	var wc2 WriterCloser = NewBufferedWriterCloser()
	wc2.Write([]byte("Hello again"))
	wc2.Close()

	bwc := wc2.(*BufferedWriterCloser)  // Example of Type Conversion
	fmt.Println(bwc)

	// wwc := wc2.(io.Reader)  // Example of failed Type conversion: Missing method Read. Panics!!!
	// fmt.Println(wwc)

	// Comma-ok syntax. Does not panic
	r, ok := wc2.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

	// Empty interface
	var myObj interface{} = NewBufferedWriterCloser()

	// Type-cast into a WriterCloser using comma-ok syntax
	if wc3, ok := myObj.(WriterCloser); ok {
		wc3.Write([]byte("Yay, typecast worked"))
		wc3.Close()
	}
	r2, ok := myObj.(io.Reader)
	if ok {
		fmt.Println(r2)
	} else {
		fmt.Println("Conversion failed")
	}

	// Type Switch.  dot then parentheses with type variable
	var i2 interface{} = 0
	switch i2.(type) {
	case int:
		fmt.Println("i2 is an integer")
	case string:
		fmt.Println("i2 is a string")
	default:
		fmt.Println("I2 don't know what i2 is")
	}

	var wc4 WriterCloser = myWriterCloser{}
	fmt.Println(wc4)
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

// Interface composition example

// Closer interface
type Closer interface {
	Close() error
}

// WriterCloser interface
type WriterCloser interface {
	Writer
	Closer
}

type myWriterCloser struct {}

// func (mwc *myWriterCloser) Write(data []byte) (int, error) { This would be an error
// - When working with types, method set is all the methods regardless of receiver type
// - When working with interfaces and using the concrete value (not the address), the method set for a value is any
// method that has a value as the receiver
// Can fix by using `var wc4 WriterCloser = &myWriterCloser{}`
// - The method set for a pointer is the sum of the value and pointer receiver methods.
func (mwc myWriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

func (mwc myWriterCloser) Close() error {
	return nil
}

// BufferedWriterCloser struct
type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

// Close method
func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// NewBufferedWriterCloser constructor
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}