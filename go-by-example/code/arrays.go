package gobyex

import "fmt"

func Array() {
	var a [5]int
	fmt.Println(a)

	//a[6] = 9:error

	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	c := [5]int{2, 4, 6, 8, 0}
	fmt.Println(c)

	var twoD [2][3]int

	for i := 1; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			twoD[i-1][j-1] = i + j
		}
	}
	fmt.Println(twoD)

}
