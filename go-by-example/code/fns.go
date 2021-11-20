package gobyex

import "fmt"

func Fns() {
	fmt.Println(sum(1, 2, 3, 4))
	a := []int{1, 2, 3, 4}
	fmt.Println(sum(a...))
}

func sum(nums ...int) int {

	s := 0
	for _, v := range nums {
		s += v
	}

	return s
}
