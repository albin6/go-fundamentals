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

	d := [5]int{}

	fmt.Println("empty arr =>", d)

	for i := range 5 {
		d[i] = i + 1
	}

	fmt.Println("after filling =>", d)

	// 2D array
	twoDArr := [2][2]int{
		{1, 2},
		{4, 5},
	}

	fmt.Println(twoDArr)

	sum := 0

	for i := 0; i < len(twoDArr); i++ {
		for j := 0; j < len(twoDArr[i]); j++ {
			sum += twoDArr[i][j]
		}
	}

	fmt.Println("sum of two d array elements", sum)
}
