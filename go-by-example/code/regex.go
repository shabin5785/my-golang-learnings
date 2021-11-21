package gobyex

import (
	"fmt"
	"regexp"
)

func Regex() {

	m, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(m)

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peac, peach , punch"))

	fmt.Println(r.FindStringSubmatch("peach punch"))

	//safer global regex

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regex ", r)

	fmt.Println(r.ReplaceAllString("a peach", "fruit"))
}
