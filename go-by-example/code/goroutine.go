package gobyex

import (
	"fmt"
	"time"
)

func GoRoutine() {

	f("word")
	go f("hello")
	fmt.Println("after")

	go func(msg string) {
		fmt.Println(msg)
	}("anonymous")

	time.Sleep(time.Second * 2)
}

func f(from string) {
	for i := 0; i <= 3; i++ {
		fmt.Println(from, ":", i)
	}
}
