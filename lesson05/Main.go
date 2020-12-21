package main

import (
	"fmt"
)

// Example of package-level constant being shadowed
const a int16 = 27

func main() {
	fmt.Println("Lesson 05: Constants")
	fmt.Println("- Naming convention")
	fmt.Println("- Typed constants")
	fmt.Println("- Untyped constants")
	fmt.Println("- Enumerated constants")
	fmt.Println("- Enumeration expressions")

	// typed constant that is not exported
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	// Not allowed because not assignable at compile-time
	// const myConst2 float64 = math.Sin(1.57)

	const a int = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

	// constants and variables of the same type can be added together
	const q int = 14
	var r int = 16
	fmt.Printf("%v, %T\n", q+r, q+r)

	// untyped constants -- implict conversions
	const untypedConst1 = 42 // by itself would be 42, int.  seen by compiler as a literal 42
	var typedConst1 int16 = 27
	fmt.Printf("%v, %T\n", untypedConst1 + typedConst1, untypedConst1 + typedConst1)  // 69, int16 (42 + typedConst1)

}
