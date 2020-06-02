/* =====================================================================
 * -- Name ------ : Assing variable to anonymous functions
 * -- Date ------ : Jun 2, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : Assign a func to a variable, then call that func
 ===================================================================== */

package main

import "fmt"

func main() {
	sayMyName := func(name string) {
		fmt.Println("Your name is", name)
	}

	sayMyName("Carlos")
}
