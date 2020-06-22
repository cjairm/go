/* =====================================================================
 * -- Name ------ : For infinite
 * -- Date ------ : May 8, 2020
 * -- Author ---- : Carlos Mendez
 ===================================================================== */

package main

import "fmt"

func main() {
	bd := 1991
	myAge := 0
	for {
		if bd > 2020 {
			break
		}
		fmt.Println(bd)
		fmt.Println(myAge)
		bd++
		myAge++
	}
}
