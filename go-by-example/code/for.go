package gobyex

import "fmt"

func For() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 5; j++ {
		if j%3 == 0 {
			continue
		}
		fmt.Println(j)

		if j == 5 {
			break

		}
	}
}
