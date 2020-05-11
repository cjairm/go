/* =====================================================================
 * -- Name ------ : Operators
 * -- Date ------ : May 8, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : Testing operators in order to check how they work
 ===================================================================== */

package main

import "fmt"

func main() {
	equalTo := (3 == 3)
	lessOrEqual := (3 <= 0)
	greaterOrEqual := (3 >= 0)
	differentThan := (3 != 3)
	lessThan := (0 < 3)
	greaterThan := (0 > 3)

	fmt.Println("equal to: ", equalTo)
	fmt.Println("less or equal to: ", lessOrEqual)
	fmt.Println("greater or equal to: ", greaterOrEqual)
	fmt.Println("different than: ", differentThan)
	fmt.Println("less than: ", lessThan)
	fmt.Println("greater than: ", greaterThan)
}
