package main

import (
	"fmt"
	"runtime"
)

func main() {
	ch := make(chan int)

	fmt.Println("Goroutines:", runtime.NumGoroutine())

	fs := func() {
		for j := 0; j < 10; j++ {
			ch <- j
		}
		close(ch)
	}

	for i := 0; i < 10; i++ {
		go fs()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	fmt.Println("Goroutines:", runtime.NumGoroutine())

	for k := 0; k < 99; k++ {
		fmt.Println(k, <-ch)
	}
}
