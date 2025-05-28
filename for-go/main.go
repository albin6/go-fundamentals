package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	fmt.Print("----------\n")

	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	fmt.Print("----------\n")

	for {
		fmt.Println("hello")
		break
	}

	fmt.Print("----------\n")

	for i := range 3 {
		fmt.Println("range", i)
	}

	fmt.Print("----------\n")

	for j := range 5 {
		if j%2 == 0 {
			continue
		}

		fmt.Println(j)
	}

}
