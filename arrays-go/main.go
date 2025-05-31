package main

import "fmt"

func main() {
	// array is a numbered sequence of elements of a specific length
	fmt.Println("Arrays in Go")

	var a [5]int
	a[0] = 10

	fmt.Println(a)

	fmt.Println(len(a))

	x := [3]int{3, 2, 1}

	fmt.Println(x)

	b := [...]int{100, 200, 300}

	fmt.Println(b)

	c := [...]int{2, 2: 2, 4: 2}

	fmt.Println(c)
}
