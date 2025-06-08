package main

import "fmt"

func sumOfEven(nums [5]int) int {
	sum := 0
	for _, num := range nums {
		if num%2 == 0 {
			sum += num
		}
	}

	return sum
}

func main() {
	nums := [5]int{1, 2, 3, 4, 5}

	var sum int = sumOfEven(nums)

	fmt.Println("sum of even:", sum)
}
