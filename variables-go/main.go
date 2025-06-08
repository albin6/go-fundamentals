package main

import "fmt"

const name string = "Albin"

func main() {
	var x = 10
	y := 20
	fmt.Println(x + y)

	fmt.Println(name)

	const name = "Arshad"

	fmt.Println(name)

	const num1, num2 int = 1, 2

	fmt.Println(num1 + num2)
}
