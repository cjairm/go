/* =====================================================================
 * -- Name ------ : Array
 * -- Date ------ : May 10, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : Using composite literal will hold an array with N values
 ===================================================================== */

package main

import "fmt"

func main() {
	var myArray [5]int
	myArray[0] = 100
	myArray[1] = 200
	myArray[2] = 300
	myArray[3] = 400
	myArray[4] = 500
	fmt.Printf("%T\n", myArray)
	fmt.Printf("%v\n", myArray)

	for i, ele := range myArray {
		fmt.Printf("The position is: %v \t the element value is: %v \n", i, ele)
	}
}
