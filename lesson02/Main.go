package main

import "fmt"

// Package level variable declarations
// var actorName string = "Elisabeth Sladen"
// var companion string = "Sarah Jane Smith"
// var doctorName int = 3
// var season int = 11

// Grouping var declarations
var (
 actorName string = "Elisabeth Sladen"
 companion string = "Sarah Jane Smith"
 doctorName int = 3
 season int = 11
)

func main() {
	fmt.Printf("%s\n", actorName )
}