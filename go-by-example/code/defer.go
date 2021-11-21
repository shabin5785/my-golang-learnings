package gobyex

import (
	"fmt"
	"os"
)

func DeferFn() {

	f := createFile("defer.txt")
	defer closeFile(f)
	writeFile(f)

}

func createFile(f string) *os.File {
	fmt.Println("creating ", f)
	file, err := os.Create(f)
	if err != nil {
		panic(err)
	}
	return file
}

func closeFile(f *os.File) {
	fmt.Println("closing ", f.Name())
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error : %v\n", err)
		os.Exit(1)
	}
}

func writeFile(f *os.File) {
	fmt.Println("Writing")
	fmt.Fprintf(f, "Some data")
}
