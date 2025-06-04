package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {
	sum := sum(10, 20)

	fmt.Println("Sum of 10 and 20 =>", sum)
}
