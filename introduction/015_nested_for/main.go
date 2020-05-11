/* =====================================================================
 * -- Name ------ : For nested
 * -- Date ------ : May 8, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : First for loop in a loop
 ===================================================================== */

package main

import "fmt"

func main() {
	for i := 10; i < 51; i++ {
		fmt.Printf("%v \n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U \n", i)
		}
	}
}
