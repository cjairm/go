/* =====================================================================
 * -- Name ------ : For conditional
 * -- Date ------ : May 8, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : Iterate a for until a condition
 ===================================================================== */

package main

import "fmt"

func main() {
	bd := 1991
	myAge := 0
	for bd <= 2020 {
		fmt.Println(bd)
		fmt.Println("My age: ", myAge)
		myAge++
		bd++
	}
}
