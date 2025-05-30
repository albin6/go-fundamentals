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
	// show how the case expressions can be non-constants.
	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() > 12:
		fmt.Println("This is after noon")
	default:
		fmt.Println("This is before noon")
	}

	// A type switch compares types instead of values.
	// You can use this to discover the type of an interface value.
	// In this example, the variable t will have the type corresponding to its clause.
	whoAmI := func(i interface{}) {
		switch j := i.(type) {
		case bool:
			fmt.Println("this is a boolean")
		case int:
			fmt.Println("this is an interger")
		default:
			fmt.Printf("type is unknown %T\n", j)
		}
	}

	whoAmI(true)
	whoAmI(1)
	whoAmI("hello")
}
