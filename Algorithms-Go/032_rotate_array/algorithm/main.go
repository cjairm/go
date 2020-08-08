// Package algorithm contains different algorithms
package algorithm

// RotateArray rotate the array
func RotateArray(xi []int, k int) []int {
	if k <= 0 {
		return xi
	}
	lastPos := len(xi) - 1
	lastEl := xi[lastPos]
	xxi := []int{}
	xxi = append(xxi, lastEl)
	for _, v := range xi[0:lastPos] {
		xxi = append(xxi, v)
	}
	return RotateArray(xxi, k-1)
}
