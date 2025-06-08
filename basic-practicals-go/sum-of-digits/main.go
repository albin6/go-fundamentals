package main

import "fmt"

func sumOfDigit(num int) int {
	if num == 0 {
		return 0
	}
	return num%10 + sumOfDigit(num/10)
}

func main() {
	var num, sum int
	fmt.Print("Enter the number :")
	fmt.Scan(&num)

	sum = sumOfDigit(num)
	fmt.Println("Sum of digits: ", sum)
}
