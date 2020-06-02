/* =====================================================================
 * -- Name ------ : Functions with params
 * -- Date ------ : Jun 2, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description :
	● create a func with the identifier foo that
		○ takes in a variadic parameter of type int
		○ pass in a value of type []int into your func (unfurl the []int)
		○ returns the sum of all values of type int passed in
	● create a func with the identifier bar that
		○ takes in a parameter of type []int
		○ returns the sum of all values of type int passed in
 ===================================================================== */

package main

import "fmt"

func main() {
	numsToSum := []int{1, 2, 10}
	sum := sumNums(numsToSum...)
	fmt.Println(sum)

	numsToSubs := []int{10, 5, 1}
	subs := subsNums(numsToSubs)
	fmt.Println(subs)
}

func sumNums(xi ...int) int {
	// fmt.Printf("Type: %T, Value: ", xi)
	var totalSum int
	for _, val := range xi {
		totalSum += val
	}
	return totalSum
}

func subsNums(xi []int) int {
	// fmt.Printf("Type: %T, Value: ", xi)
	var totalSub int
	if len(xi) <= 0 {
		return 0
	}
	totalSub = xi[:1][0]
	for _, val := range xi[1:] {
		totalSub -= val
	}
	return totalSub
}
