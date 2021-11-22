package gobyex

import (
	"bufio"
	"fmt"
	"net/http"
)

func HttpFn() {

	resp, err := http.Get("http://gobyexample.com")
	check(err)

	defer resp.Body.Close()
	fmt.Println(resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; i < 5 && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
