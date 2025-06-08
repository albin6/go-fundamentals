package main

import "fmt"

func printTable(num int) {
	for i := range 10 {
		fmt.Println(i+1, " x ", num, " = ", (i+1)*num)
	}
}

func main() {
	var num int

	fmt.Print("Enter the number : ")
	fmt.Scan(&num)

	printTable(num)
}
