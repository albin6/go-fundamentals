package main

import "fmt"

func main() {
	const s = "สวัสดี"

	for i := range len(s) {
		fmt.Printf("%x", s[i])
	}
}
