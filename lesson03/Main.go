package main

import (
	"fmt"
)

// Naming variables.
// - controls visibility
// - naming conventions

// This lowercase 'i' variable is scoped to this package
var i int = 27

// SeasonChapterName 
// This is exported to outside this package
// Variable name should be verbose but with in reason
var SeasonChapterName string = "Season 2 Chapter 5"

func main() {
	// This prints out the package-level variable
	fmt.Println(i)

	// The following variable declaration shadows the package-level variable
	var i int = 42
	fmt.Println(i)

  // The following would be a compile-time error because 'j' is not used
	// j := 13

	// Short variable name length should indicate the short-use of the variable
	// Longer variable names should indicate longer-use
	var seasonNumber int = 11
	fmt.Println(seasonNumber)

	// Acronyms in variable names. Best practice is to keep 'URL' uppercase
	var theURL string = "https://www.example.com":
}