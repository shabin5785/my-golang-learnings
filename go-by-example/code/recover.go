package gobyex

import "fmt"

func thisPanics() {
	panic("something is wrong")
}

func RecoverFn() {

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("recoverd ", r)
		}
	}()

	thisPanics()
	fmt.Println("wont reac here as it gets to recover")

}
