package gobyex

import "fmt"

func Pointers() {

	a := 1
	simplePointer(&a)
	fmt.Println(a)

}

func simplePointer(sptr *int) {
	fmt.Println(sptr)
	*sptr++
}
