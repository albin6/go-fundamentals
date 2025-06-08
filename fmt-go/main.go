package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter you name :")
	fmt.Scan(&name)             // Albin Aji (only reads first word)
	fmt.Println("Hello ", name) // Hello  Albin

	var firstName string
	var lastName string
	fmt.Println("Enter your first and last name") // Ends with newline
	fmt.Scanln(&firstName, &lastName)             // Albin Aji
	fmt.Println("Hello", firstName+" "+lastName)  // Hello Albin Aji

	var firstNameOnly string
	var age int
	fmt.Printf("Enter your firstNameOnly and age :")
	fmt.Scanf("%s %d", &firstNameOnly, &age)
	fmt.Printf("Entered firstNameOnly %s and age %d", firstNameOnly, age)
}
