package main

import "fmt"

func main() {
	fmt.Println("slices")
	var s []string

	fmt.Println("slice s =>", s, "eq nill =>", s == nil, "length =>", len(s))

	x := make([]int, 3)

	fmt.Println("using make", x)
}
