package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func multipleReturn() (int, int) {
	return 2, 4
}

func variadic(nums ...int) {
	for num := range nums {
		fmt.Println(num)
	}
}

func main() {
	sum := sum(10, 20)

	fmt.Println("Sum of 10 and 20 =>", sum)

	a, b := multipleReturn()

	fmt.Println(a, b)

	_, y := multipleReturn()

	fmt.Println(y)

	fmt.Println("Variadic function example:")

	arr := []int{1, 2, 3, 4, 5}

	variadic(arr...)
}
