/* =====================================================================
 * -- Name ------ : Raw String
 * -- Date ------ : May 8, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : This obey each character inserted included enters for example....
 ===================================================================== */

package main

import "fmt"

func main() {
	myString := `here is something
	as 
	a 
	raw string
	literal
	"including doble quotes"
	nice`
	fmt.Println(myString)
}
