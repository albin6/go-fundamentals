package main

import "fmt"

func findLargest(a int, b int, c int) int {
	var largest int
	if a > b {
		largest = max(a, c)
	} else {
		largest = max(b, c)
	}

	return largest
}

func main() {
	var largest int = findLargest(2, 1, 3)

	fmt.Println("Largest number is :", largest)
}
