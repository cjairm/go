/* =====================================================================
 * -- Name ------ : Wait Group
 * -- Date ------ : Jun 21, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : This exercise will reinforce our understanding of method sets:
	● create a type person struct
	● attach a method speak to type person using a pointer receiver
		○ *person
	● create a type human interface
		○ to implicitly implement the interface, a human must have the speak method
	● create func “saySomething”
		○ have it take in a human as a parameter
		○ have it call the speak method
	● show the following in your code
		○ you CAN pass a value of type *person into saySomething
		○ you CANNOT pass a value of type person into saySomething
	● here is a hint if you need some help
		○ https://play.golang.org/p/FAwcQbNtMG

	Receivers  Values
	-----------------------------------------------
	(t T)      T and *T
	(t *T)     *T
 ===================================================================== */

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
