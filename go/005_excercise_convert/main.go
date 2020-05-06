/* =====================================================================
 * -- Name ------ : Convert
 * -- Date ------ : May 6, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : The word casting does not exists in Go. The word used insted is conversions.
 *					Convert one variable to assing a value to another variable
 ===================================================================== */

package main

import "fmt"

type mytype int

var x mytype
var y int

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 100
	fmt.Printf("%v\n", x)
	y = int(x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)
}
