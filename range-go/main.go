package main

import "fmt"

func main() {
	nums := []int{4, 3, 2, 1}

	for i, num := range nums {
		fmt.Println(i, num)
	}

	// iterating over map key/value pairs

	kvp := map[string]any{"name": "Albin", "age": 21}

	for key, value := range kvp {
		fmt.Println(key, value)
	}

	str := "golang"

	for idx, char := range str {
		fmt.Println(idx, char)
	}
}
