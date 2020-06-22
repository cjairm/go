package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname  string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello, my name is " + p.firstname + " " + p.lastname + ".")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		firstname: "Carlos",
		lastname:  "Mendez",
	}

	saySomething(&p1)
}
