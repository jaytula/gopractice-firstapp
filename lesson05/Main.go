package main

import (
	"fmt"
)

// Example of package-level constant being shadowed
const a int16 = 27

// enumerated constants. iota is scoped to the const block
const (
	stuffA = iota
	stuffB
	stuffC
)

const (
	_ = iota + 5  // Discard initial generated value
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

// bytes
const (
	_ = iota                  // ignore first value
	KB = 1 << (10 * iota)     // 1024 Bytes
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

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
	fmt.Printf("%v, %T\n", untypedConst1+typedConst1, untypedConst1+typedConst1) // 69, int16 (42 + typedConst1)

	// enumerated constants
	fmt.Printf("%v, %T\n", stuffA, stuffA) // (0, int)
	fmt.Printf("%v, %T\n", stuffB, stuffB) // (1, int)
	fmt.Printf("%v, %T\n", stuffC, stuffC) // (2, int)

	var specialistType int = catSpecialist
	fmt.Printf("%v\n", catSpecialist)
	fmt.Printf("%v\n", specialistType == catSpecialist)

	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)
}
