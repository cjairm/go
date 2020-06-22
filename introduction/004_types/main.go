/* =====================================================================
 * -- Name ------ : Own types
 * -- Date ------ : May 6, 2020
 * -- Author ---- : Carlos Mendez
 ===================================================================== */

package main

import "fmt"

type mytype int

var x mytype

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 100
	fmt.Printf("%v\n", x)
}
