package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func changeMe(p *person) {
	(*p).age = 89
	// p.age = 89
}

func main() {
	p := person{
		firstName: "Carlos",
		lastName:  "Mendez",
		age:       33,
	}

	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}
