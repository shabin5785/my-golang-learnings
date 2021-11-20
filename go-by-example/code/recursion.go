package gobyex

import "fmt"

func Recursion() {
	fmt.Println(fact(10))
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
