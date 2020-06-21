/* =====================================================================
 * -- Name ------ : Wait Group
 * -- Date ------ : Jun 18, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description :
	● In addition to the main goroutine, launch two additional goroutines
		○ Each additional goroutine should print something out
	● Use waitgroups to make sure each goroutine finishes before your program exists
 ===================================================================== */

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	maxNumber := 100
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	go printEvenNumbers(maxNumber)
	go printOddNumbers(maxNumber)
	printAllNumbers(maxNumber)

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Wait()
}

func printEvenNumbers(mn int) {
	for i := 0; i < mn; i++ {
		if i%2 != 0 {
			fmt.Println("I'm an EVEN number: ", i)
		}
	}
	wg.Done()
}

func printOddNumbers(mn int) {
	for i := 0; i < mn; i++ {
		if i%2 == 0 {
			fmt.Println("I'm an ODD number: ", i)
		}
	}
	wg.Done()
}

func printAllNumbers(mn int) {
	for i := 0; i < mn; i++ {
		fmt.Println("I'm NUMBER: ", i)
	}
}
