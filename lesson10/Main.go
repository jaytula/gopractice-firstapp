package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Control Flow
// - Defer - Invoke function but delay execution
// - Panic - How runtime can trigger and how to manually trigger
// - Recover - How to recover from panic

func deferExample1() {
	fmt.Println("deferExample1: start")
	defer fmt.Println("deferExample1: middle")  // Moves to after the main function. But before it returns
	fmt.Println("deferExample1: end")
}

// Demonstrates that deferred functions executed in LIFO order: end, middle, start
func deferExample2() {
	defer fmt.Println("deferExample2: start")
	defer fmt.Println("deferExample2: middle")  // Moves to after the main function. But before it returns
	defer fmt.Println("deferExample2: end")
}

// Using defer to close resource after its been opened and checked for errors
func httpExample() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func main() {
	fmt.Println("main: start")
	deferExample1()
	deferExample2()
	httpExample()
	fmt.Println("main: end")
}

