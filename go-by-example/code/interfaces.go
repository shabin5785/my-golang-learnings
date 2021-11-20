package gobyex

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	peri() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) peri() float64 {
	return r.height + r.width
}
func (c circle) peri() float64 {
	return 2 * math.Pi * c.radius
}

func measure(s shape) {
	fmt.Println(s.area())
	fmt.Println(s.peri())
}

func Interfaces() {
	r := rectangle{4, 5}
	measure(r)

	c := circle{4}
	measure(c)
}
