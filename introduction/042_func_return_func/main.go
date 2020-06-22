/* =====================================================================
 * -- Name ------ : 
 * -- Date ------ : 
 * -- Author ---- : Carlos Mendez
 * -- Description :
	
 ===================================================================== */

package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	newFunc := foo()
	fmt.Println("Result: ", newFunc())
}

func foo() func() int {
	i := 7
	return func() int {
		return i
	}
}
