package main

import "fmt"

func alterValue(val int) {
	val = 0
}

func alterValuePtr(ptr *int) {
	*ptr = 0
}

func main() {
	num := 1
	fmt.Println("num before calling altervalue =>", num)

	alterValue(num)
	fmt.Println("num after calling altervalue =>", num)

	alterValuePtr(&num)
	fmt.Println("num after calling altervaluePtr =>", num)

	fmt.Println("pointer =>", &num)

}
