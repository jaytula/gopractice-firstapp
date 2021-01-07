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
}