/* =====================================================================
 * -- Name ------ : For - module
 * -- Date ------ : May 8, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : found module of 4 in numbers 4-90
 ===================================================================== */

package main

import "fmt"

func main() {
	fmt.Println("===Module 4===")
	for i := 4; i <= 90; i++ {
		fmt.Printf("Module %v of %v \n", i%4, i)
	}
}
