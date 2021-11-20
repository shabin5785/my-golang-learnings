package gobyex

import "fmt"

func Closure() {

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
