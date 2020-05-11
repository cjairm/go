/* =====================================================================
 * -- Name ------ : Slice
 * -- Date ------ : May 10, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : Using composite literal creating slice type
 ===================================================================== */

package main

import "fmt"

func main() {
	mySlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%T\n", mySlice)
	fmt.Printf("%v\n", mySlice)

	for i, ele := range mySlice {
		fmt.Printf("The position is: %v \t the element value is: %v \n", i, ele)
	}
}
