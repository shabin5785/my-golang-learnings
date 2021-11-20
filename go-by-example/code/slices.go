package gobyex

import "fmt"

func Slice() {

	s := make([]string, 3)
	fmt.Println(s)
	s[0] = "a"
	fmt.Println(s)

	s = append(s, "asd")
	fmt.Println(s)

	s = append(s, "A", "B")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	c[0] = "Z"
	fmt.Println(c)
	fmt.Println(s)

	fmt.Println(c[1:2])

	t := []string{"a", "b", "c"}
	fmt.Println(t)

	//slices can have varying lengh of inner slices, unlike array
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)
}
