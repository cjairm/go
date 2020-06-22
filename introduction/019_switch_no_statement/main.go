/* =====================================================================
 * -- Name ------ : switch
 * -- Date ------ : May 8, 2020
 * -- Author ---- : Carlos Mendez
 ===================================================================== */

package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("this is printed")
	case false:
		fmt.Println("this is NOT printed")
	}
}
