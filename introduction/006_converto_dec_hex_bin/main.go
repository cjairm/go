/* =====================================================================
 * -- Name ------ : Convert to
 * -- Date ------ : May 8, 2020
 * -- Author ---- : Carlos Mendez
 ===================================================================== */

package main

import "fmt"

var myNumber int

func main() {
	myNumber = 100
	fmt.Printf("My decimal number is: %d\n", myNumber)
	fmt.Printf("My binary number is: %b\n", myNumber)
	fmt.Printf("My hexadecimal number is: %X\n", myNumber)
}
