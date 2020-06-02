/* =====================================================================
 * -- Name ------ : Functions attach method to type
 * -- Date ------ : Jun 2, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description :
	● Create a user defined struct with
		○ the identifier “person”
		○ the fields:
			■ first
			■ last
			■ age
	● attach a method to type person with
		○ the identifier “speak”
		○ the method should have the person say their name and age
	● create a value of type person
	● call the method from the value of type person
 ===================================================================== */

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
