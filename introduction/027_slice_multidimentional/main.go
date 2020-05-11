/* =====================================================================
 * -- Name ------ : States multidimentional
 * -- Date ------ : May 10, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : Slice that store many slices inside
 ===================================================================== */

package main

import "fmt"

func main() {
	firstSlice := []string{"James", "Bond", "Shaken, not stirred"}
	secondSlice := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	myMulti := [][]string{firstSlice, secondSlice, firstSlice}
	fmt.Println(myMulti)

	for _, multiEle := range myMulti {
		fmt.Println(multiEle)
		for _, singleEle := range multiEle {
			fmt.Printf("\t %v \n", singleEle)
		}
	}
}
