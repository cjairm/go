/* =====================================================================
 * -- Name ------ : if
 * -- Date ------ : May 8, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : working for first time with if statement
 ===================================================================== */

package main

import "fmt"

func main() {
	if true {
		fmt.Println("This is printed")
	}

	if false {
		fmt.Println("This is NOT printed")
	}

	if 2 == 2 {
		fmt.Println("This is printed because 2 is equal 2")
	}

	if 2 == 22 {
		fmt.Println("This is NOT printed bacause 2 is diferent than 22")
	}
}
