package gobyex

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FilesProcessFn() {

	f, _ := os.Open("defer.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		u := strings.ToUpper(scanner.Text())
		fmt.Println(u)
	}

	filepaths()
}

func filepaths() {
	p := filepath.Join("dir1", "dir2", "f.txt")
	fmt.Println(p)

	directories()
}

func directories() {
	err := os.Mkdir("sub", 0755)
	check(err)

	defer os.RemoveAll("sub")

	tempFD()

}

func tempFD() {

	f, err := os.CreateTemp("", "sample")
	check(err)

	fmt.Println("temp ", f.Name())

	_, err = f.WriteString("hello there")
	check(err)

	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("dir ", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "f1")
	os.WriteFile(fname, []byte("wrote"), 0666)

}
