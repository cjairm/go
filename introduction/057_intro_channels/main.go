package main

import "fmt"

func main() {
	// c := make(chan int, 1)
	c := make(chan int)

	f := func() {
		c <- 42
	}

	go f()

	fmt.Println(<-c)
}
