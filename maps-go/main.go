package main

import "fmt"

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
}
