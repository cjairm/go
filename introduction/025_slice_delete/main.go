/* =====================================================================
 * -- Name ------ : Slice delete
 * -- Date ------ : May 10, 2020
 * -- Author ---- : Carlos Mendez
 ===================================================================== */

package main

import "fmt"

func main() {
	mySlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	newSlice := append(mySlice[:3], mySlice[6:]...)

	fmt.Println(newSlice)
}
