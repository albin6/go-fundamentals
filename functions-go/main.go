package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func multipleReturn() (int, int) {
	return 2, 4
}

func main() {
	sum := sum(10, 20)

	fmt.Println("Sum of 10 and 20 =>", sum)

	a, b := multipleReturn()

	fmt.Println(a, b)

	_, y := multipleReturn()

	fmt.Println(y)
}
