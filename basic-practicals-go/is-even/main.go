package main

import "fmt"

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func main() {
	result := isEven(2)
	fmt.Println("Is Even :", result)
}
