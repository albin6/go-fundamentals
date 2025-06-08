package main

import "fmt"

func main() {
	// opiton one
	var name string
	fmt.Print("Enter first name :")
	fmt.Scan(&name)

	fmt.Println("first name ", name)

	fmt.Println("Enter full name :")
	fmt.Scanf("%s", &name)

	fmt.Println("Full name :", name)

}
