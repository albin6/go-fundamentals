package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("slices")
	var s []string

	fmt.Println("slice s =>", s, "eq nill =>", s == nil, "length =>", len(s))

	x := make([]int, 3)

	fmt.Println("using make", x)

	x[0] = 1
	x[1] = 2
	x[2] = 3

	c := make([]int, len(x))

	copy(c, x)

	fmt.Println("After making copy =>", c)

	// appending values
	fmt.Println("Before appending => x =>", x)

	x = append(x, 4)

	fmt.Println("After appending 4 => x =>", x)

	x = append(x, 5, 6)

	fmt.Println("After appending 5, 6 => x =>", x)

	// slice operation in with slices
	l := x[1:5]

	fmt.Println("slicing from 1 : 5 =>", l) // index 1 to 4 (5 is excluding)

	m := x[1:]

	fmt.Println("slicing from 1 : =>", m) // index 1 to end (1 includes)

	n := x[:5]

	fmt.Println("slicing till : 5 =>", n) // index 0 to 4 (5 is excluding)

	t1 := []int{1, 2, 3, 4}
	t2 := []int{1, 2, 3, 4}

	if slices.Equal(t1, t2) {
		fmt.Println("t1 == t2")
	}

	twoD := make([][]int, 3)

	for i := range len(twoD) {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("two d slice =>", twoD)
}
