package gobyex

import (
	"fmt"
	"os"
)

func Env() {
	for _, v := range os.Environ() {
		fmt.Println(v)
	}
}
