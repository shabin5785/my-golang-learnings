package gobyex

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func FilesFn() {

	p := fmt.Println
	dat, err := os.ReadFile("defer.txt")
	check(err)
	p(string(dat))

	f, err := os.Open("defer.txt")
	check(err)

	//read few bytes
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	p(string(b1[:n1]))

	f.Close()

	writeFileFn()
}

func writeFileFn() {

	err := os.WriteFile("defer.txt", []byte("hello there"), 0644)
	check(err)

	f, err := os.Create("defer.txt")
	check(err)

	w := bufio.NewWriter(f)
	w.WriteString("buffered")
	w.Flush()
	f.Close()
}
