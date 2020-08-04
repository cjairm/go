// Package algorithm contains different algorithms
package algorithm

// ProductOfArray get the product of all elements in the array
func ProductOfArray(xi []int) int {
	if len(xi) == 0 {
		return 0
	} else if len(xi) == 1 {
		return xi[0]
	}

	return xi[0] * ProductOfArray(xi[1:len(xi)])
}
