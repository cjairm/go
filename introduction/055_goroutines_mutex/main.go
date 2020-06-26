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
	var m sync.Mutex

	ctr := 0
	goFunc := func() {
		m.Lock()
		v := ctr
		v++
		ctr = v
		fmt.Println(ctr)
		m.Unlock()
		wg.Done()
	}
	fmt.Println(ctr)

	for i := 0; i < mctr; i++ {
		go goFunc()
	}
	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", ctr)
}
