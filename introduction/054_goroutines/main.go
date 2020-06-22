/* =====================================================================
 * -- Name ------ : Go Routines
 * -- Date ------ : Jun 22, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description :
	● Using goroutines, create an incrementer program
		○ have a variable to hold the incrementer value
		○ launch a bunch of goroutines
			■ each goroutine should
				● read the incrementer value
					○ store it in a new variable
				● yield the processor with runtime.Gosched()
				● increment the new variable
				● write the value in the new variable back to the incrementer variable
	● use waitgroups to wait for all of your goroutines to finish
	● the above will create a race condition.
	● Prove that it is a race condition by using the -race flag
	● if you need help, here is a hint: ​ https://play.golang.org/p/FYGoflKQej
 ===================================================================== */

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
