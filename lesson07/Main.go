package main
import (
	"fmt"
)

func main() {
	fmt.Println("Maps and Structs")
	
	// Maps map a key type to a value type
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas": 27862596,
		"Florida": 20612439,
		"New York": 19745289,
		"Penssylvania": 12802503,
		"Illinois": 12801539,
		"Ohio": 11614373,
	}
	fmt.Println(statePopulations)

	// Key types are restricted to types testable for equality
	// - boolean
	// - all numeric types
	// - strings
	// - pointers
	// - interfacees
	// - structs
	// - arrays
	// - channels

	// Cannot be key types
	// - slices
	// - maps
	// - other functions

	// m := map[[]int]string{} // Invalid map key type
	n := map[[3]int]string{} // Ok
	fmt.Println(n)

	// Declaring a map with make
	statePopulations2 := make(map[string]int)
  statePopulations2 = map[string]int{
		"Ca": 12345,
		"Nv": 888,
		"Mo": 333,
	}
	fmt.Println(statePopulations2)

	// Getting an element of map
	fmt.Println(statePopulations2["Nv"])

	// Add value to map
	statePopulations2["Fl"] = 9999
	fmt.Println(statePopulations2)

	// Delete from map
	delete(statePopulations2, "Ca")
	fmt.Println(statePopulations2)

	fmt.Println(statePopulations2["Cda"])  // Returns zero for invalid key

	// Comma, ok to determine if key is valid
	pop, ok := statePopulations2["Nv"]
	fmt.Println(pop, ok)

	// Throw away value
	_, ok2 := statePopulations2["Ca"]
	fmt.Println(ok2)

	// Length of map works
	fmt.Println(len(statePopulations2))

	// Maps are reference types. Add
	statePopulations2Ref := statePopulations2
	statePopulations2["Ut"] = 27
	statePopulations2["Nv"] = 999
	fmt.Println(statePopulations2)
	fmt.Println(statePopulations2Ref)
}