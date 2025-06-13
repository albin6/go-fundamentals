package main

import "fmt"

func findDuplicates(nums []int) []int {
	duplicates := []int{}
	freq := make(map[string]int)

	for i, num := range nums {
		if freq[num] {
			freq['h'] += 1
		}
	}

	return duplicates
}

func main() {
	arr := []int{2, 1, 2, 4, 1, 5, 8, 4}
	duplicates := findDuplicates(arr)

	fmt.Println(duplicates)
}
