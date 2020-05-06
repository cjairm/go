package main

import "fmt"

func main() {
	fmt.Println("Hello")
	foo()
	fmt.Println("For loop")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("From: Foo")
}

// control flow:
// (1) sequense
// (2) loop: iterative
// (3) conditional
