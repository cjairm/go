package main

import "fmt"

func main() {
	// Creating bidirectional channel
	ch := make(chan int)

	adding := func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}

	go adding()

	for n := range ch {
		fmt.Println("curr val: ", n)
	}
}
