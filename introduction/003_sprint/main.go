/* =====================================================================
 * -- Name ------ : Sprintf
 * -- Date ------ : May 6, 2020
 * -- Author ---- : Carlos Mendez
 ===================================================================== */

package main

import "fmt"

var x int = 100
var y string = "Carlos Mendez"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
