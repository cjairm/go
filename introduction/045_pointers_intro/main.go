/* =====================================================================
 * -- Name ------ : Pointer intro
 * -- Date ------ : Jun 5, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description :
	● Create a value and assign it to a variable.
	● Print the address of that value.
 ===================================================================== */

package main

import "fmt"

func main() {
	v := "Hello world"
	fmt.Println(`The address in memory of "`+v+`" is: `, &v)
}
