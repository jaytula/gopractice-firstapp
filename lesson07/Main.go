package main

import (
	"fmt"
)

// Doctor Struct is exported from package because it is capitalized
// Number field is exported but actorName and companions are not
type Doctor struct {
	Number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly bool
}

func main() {
	fmt.Println("Maps and Structs")

	// Maps map a key type to a value type
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Penssylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
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

	fmt.Println(statePopulations2["Cda"]) // Returns zero for invalid key

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

	// Struct Collection type - gathers information that can have different types.
	aDoctor := Doctor{
		Number:     3,
		actorName:  "Jon Pertwee",
		companions: []string{"Liz Shaw", "Jo Grant", "Sarah Jane Smith"},
	}
	fmt.Println(aDoctor)

	// Use dot-syntax to get a Struct member
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])

	// Position syntax with Struct
	// - This works but is not recommended because if struct changes, it breaks
	bDoctor := Doctor{
		3,
		"Jon Pertwee",
		[]string{"Liz Shaw", "Jo Grant", "Sarah Jane Smith"},
	}
	fmt.Println(bDoctor)

	// Example of anonymous struct
	// Use when struct is only used one time
	cDoctor := struct{ name string }{name: "John Pertwee"}
	fmt.Println(cDoctor)

	// Structs are values types
	anotherDoc := cDoctor
	cDoctor.name = "Bluewings"
	fmt.Println(cDoctor)    // {Bluewings}
	fmt.Println(anotherDoc) // {John Pertwee}

	// We can point to the same underlying data of struct with '&'
	sameDoc := &cDoctor
	cDoctor.name = "Odo"
	fmt.Println(cDoctor) // {Odo}
	fmt.Println(sameDoc) // &{Odo}

	// Embedding/Composition
	// Bird has Animal-like characteristics
	bird := Bird{}
	bird.Name = "Emu"
	bird.Origin = "Australia"
	bird.SpeedKPH = 48
	bird.CanFly = false
	fmt.Println(bird)

	// Literal syntax - need to be aware of embedding
	bird2 := Bird{Animal: Animal{Name: "Bayou", Origin: "New Zealand"}, SpeedKPH: 15, CanFly: true}
	fmt.Println(bird2)

	// - When modeling behavior, embedding is not the right choice
	// - Interfaces are more appropriate to describe common behavior
}
