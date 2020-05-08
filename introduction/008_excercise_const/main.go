/* =====================================================================
 * -- Name ------ : Type and untype constants
 * -- Date ------ : May 8, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : We'll practice how to declare constants that are type and untyped
 ===================================================================== */

package main

import "fmt"

const (
	untyped     = 10
	typed   int = 12
)

func main() {
	fmt.Println(untyped)
	fmt.Println(typed)
}
