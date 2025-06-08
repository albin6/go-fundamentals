package main

import "fmt"

func main() {
	// using var with explicit type
	var num int = 10
	// Type is specified explicitly.
	// Initialization is optional.
	// Used when you want to be clear about the type or when declaring zero-valued variables.

	var num2 int // y is initialized to 0 (zero value of int)

	fmt.Println(num, num2)

	// Using var with type inference
	var num3 = 5
	// The type is inferred from the value.
	// Go will determine the type (in this case, int) automatically.
	// Still allows var scope flexibility (package-level, block-level).

	fmt.Println(num3)

	// Short variable declaration using :=
	num4 := 10
	// Only allowed inside functions.
	// Type is inferred.
	// Most concise and common way to declare local variables.
	// Cannot be used at the package level (outside of a function).

	fmt.Println(num4)

	// Multiple variable declaration
	var a, b, c int = 1, 2, 3
	// or
	x, y, z := 1, 2, 3
	// Allows you to declare and initialize multiple variables at once.
	// Can mix values or declare zero values.

	fmt.Println(a, b, c)
	fmt.Println(x, y, z)

	// Grouped variable declaration

	var (
		name string = "Albin"
		age  int    = 21
	)

	// Useful for organizing related variables together.
	// Improves readability especially in struct-like or config use-cases.

	fmt.Println(name, age)

	// Declaration with const
	const pi = 3.14

	// Used for constant values that do not change.
	// Type can be explicit or inferred.
	// Must be assignable at compile time
}
