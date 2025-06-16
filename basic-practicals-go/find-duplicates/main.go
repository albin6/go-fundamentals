package main

import "fmt"

func findDuplicates(nums []int) []int {
	duplicates := []int{}
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
		if freq[num] == 2 {
			// Add to duplicates only when seen the second time
			duplicates = append(duplicates, num)
		}
	}

	return duplicates
}

func main() {
	arr := []int{2, 1, 2, 4, 1, 5, 8, 4}
	duplicates := findDuplicates(arr)

	fmt.Println(duplicates)
}
