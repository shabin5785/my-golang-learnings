package gobyex

import "fmt"

func Map() {
	m := make(map[string]int)
	m["A"] = 1
	m["B"] = 2

	fmt.Println(m)

	_, ok := m["A"]
	fmt.Println(ok)

	n := map[string]int{"A": 1, "C": 2}
	fmt.Println(n)
}
