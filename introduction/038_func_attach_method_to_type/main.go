package main

import (
	"fmt"
	"strconv"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() string {
	return "Hello... My name is " + p.first + " " + p.last + " and I'm " + strconv.Itoa(p.age) + " years old"
}

func main() {
	ironman := person{
		first: "Tony",
		last:  "Stark",
		age:   42,
	}

	fmt.Println(ironman.speak())
}
