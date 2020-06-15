/* =====================================================================
 * -- Name ------ : Method Set
 * -- Date ------ : June 15, 2020
 * -- Author ---- : Todd McLeod
 * -- Description : Example
 * -- GitHub account ------ : @GoesToEleven
 ===================================================================== */

package main

import (
	"fmt"
)

func doSomething(x int) int {
	return x * 5
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)
}
