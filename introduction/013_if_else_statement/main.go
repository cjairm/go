/* =====================================================================
 * -- Name ------ : if - else if - else
 * -- Date ------ : May 8, 2020
 * -- Author ---- : Carlos Mendez
 ===================================================================== */

package main

import "fmt"

func main() {
	myAge := 29

	if myAge == 2309 {
		fmt.Println("My age is 2309")
	} else if (2020 - 1991) == 29 {
		fmt.Printf("My age is %v \n", myAge)
	} else {
		fmt.Println("No condition set as true")
	}
}
