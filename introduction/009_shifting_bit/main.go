/* =====================================================================
 * -- Name ------ : Shift bit
 * -- Date ------ : May 8, 2020
 * -- Author ---- : Carlos Mendez
 ===================================================================== */

package main

import "fmt"

func main() {
	myIntVariable := 8
	fmt.Printf("DEC %d \n", myIntVariable)
	fmt.Printf("HEX %X \n", myIntVariable)
	fmt.Printf("BIN %b \n", myIntVariable)

	myShiftedVar := myIntVariable << 1
	fmt.Printf("DEC %d \n", myShiftedVar)
	fmt.Printf("HEX %X \n", myShiftedVar)
	fmt.Printf("BIN %b \n", myShiftedVar)
}
