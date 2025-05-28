package main

import "fmt"

func main() {
	num1 := 10
	num2 := 20
	num3 := 15

	if num1 > num2 {
		if num1 > num3 {
			fmt.Println(num1, " is big")
		} else {
			fmt.Println(num3, " is big")
		}
	} else {
		if num2 > num3 {
			fmt.Println(num2, " is big")
		} else {
			fmt.Println(num3, " is big")
		}
	}

}
