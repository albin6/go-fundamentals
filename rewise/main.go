package main

import "fmt"

func main() {
	var name string = "Albin"

	fmt.Println(name)

	print_name(name)

	sum := find_sum(10, 20)

	fmt.Println("Sum is =", sum)


	grouped_declaration()
}

func print_name(name string) {
	fmt.Println("Hello ", name)
}

func find_sum(num1 int, num2 int) int {
	return  num1 + num2
}

func grouped_declaration() {
	var x, y, z = 1, 2, 3

	var (
		a int
		b string
		c float32 = 5
	)

	fmt.Println("grouped declaration =>", a, b, c)


	fmt.Println(x, y, z)
}