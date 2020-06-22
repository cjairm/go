package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	mctr := 100
	var wg sync.WaitGroup
	wg.Add(mctr)

	fmt.Println("Goroutines:", runtime.NumGoroutine())

	ctr := 0
	goFunc := func() {
		v := ctr
		runtime.Gosched()
		v++
		ctr = v
		wg.Done()
	}

	for i := 0; i < mctr; i++ {
		go goFunc()
	}
	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", ctr)
}
