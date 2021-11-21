package gobyex

import (
	"fmt"
	"sort"
)

type byLength []string

//need to implement sort interface. Len Swap Less

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func Sorting() {

	//sort method is based on type

	strs := []string{"a", "x", "d", "y"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{3, 2, 7, 1}
	sort.Ints(ints)
	fmt.Println(ints)

	fmt.Println(sort.IntsAreSorted(ints))

	//custom sort fns
	fruits := byLength{"abgv", "asd", "adsdas", "sd"}
	sort.Sort(fruits)
	fmt.Println(fruits)

}
