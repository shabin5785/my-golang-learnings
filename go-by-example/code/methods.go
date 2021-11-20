package gobyex

import "fmt"

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.height * r.width
}

func (r *rect) peri() int {
	return r.height + r.width
}

func Methods() {
	r := rect{10, 20}
	fmt.Println(r.area())
	fmt.Println(r.peri())

	rp := rect{2, 3}
	fmt.Println(rp.area())
	fmt.Println(rp.peri())
}
