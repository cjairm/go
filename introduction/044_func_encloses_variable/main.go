package main

import "fmt"

func main() {
	f := foo()
	g := foo()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	for {
		value := g()
		fmt.Println(value)
		if value >= 20 {
			break
		}
	}
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
