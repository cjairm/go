/* =====================================================================
 * -- Name ------ : Slice Piece
 * -- Date ------ : May 10, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : Using a slides create slices from the main one.
 ===================================================================== */

package main

import "fmt"

func main() {
	mySlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(mySlice[:5])
	fmt.Println(mySlice[5:])
	fmt.Println(mySlice[2:7])
	fmt.Println(mySlice[1:6])
}
