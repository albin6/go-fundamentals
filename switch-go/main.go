package main

import (
	"fmt"
	"time"
)

func main() {
	// normal switch
	i := 1
	switch i {
	case 0:
		fmt.Println("Number is 0")
		break
	}

	fmt.Println(time.Now().Weekday())

	// multiple case matching
	switch time.Now().Weekday() {
	case time.Thursday, time.Friday:
		fmt.Println("This is a friday")
	case time.Saturday, time.Sunday:
		fmt.Println("This is a weekend")
	default:
		fmt.Println("This is a weekday")
	}

	// switch like if-else
	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() > 12:
		fmt.Println("This is after noon")
	default:
		fmt.Println("This is before noon")
	}
}
