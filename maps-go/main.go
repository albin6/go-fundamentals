package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("Maps")

	m := make(map[string]int)

	m["k1"] = 10
	m["k2"] = 20

	fmt.Println(m)

	v1 := m["k1"]

	fmt.Println("v1 value =>", v1)

	fmt.Println("length of m =>", len(m))

	delete(m, "k1")

	fmt.Println("length of m after deleting k2 =>", len(m))

	// clear(m)

	// fmt.Println("m after clearing all =>", m)

	_, prs := m["k2"]

	fmt.Println("prs => ", prs)

	m1 := map[string]int{"k1": 1, "k2": 2}

	fmt.Println("m1 =>", m1)

	m2 := map[string]int{"k1": 1, "k2": 2}

	fmt.Println("m2 =>", m2)

	if maps.Equal(m1, m2) {
		fmt.Println("m1 == m2")
	}
}
