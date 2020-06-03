/* =====================================================================
 * -- Name ------ : Return function
 * -- Date ------ : Jun 3, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description :
	● Create a func which returns a func
	● assign the returned func to a variable
	● call the returned func
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
