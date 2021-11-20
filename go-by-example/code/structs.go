package gobyex

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 3
	return &p
}

func Structs() {
	fmt.Println(person{"bob", 20})
	fmt.Println(newPerson("joh"))

	sp := &person{"jh", 32}
	fmt.Println(sp.age)

	print(sp)
}

func print(p *person) {
	fmt.Println(p.name)
}
