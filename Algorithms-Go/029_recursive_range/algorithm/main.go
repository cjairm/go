// Package algorithm contains different algorithms
package algorithm

// RecursiveRange sums all the nums in the range
func RecursiveRange(i int) int {
	if i <= 1 {
		return i
	}

	return i + RecursiveRange(i-1)
}
