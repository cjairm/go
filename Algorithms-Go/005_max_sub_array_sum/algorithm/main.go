// Package algorithm contains different algorithms
package algorithm

// MaxSubArraySum look the sum of group of elements.
// Create a function called MaxSubArraySum
func MaxSubArraySum(x []int, n int) int {
	// Check array if "n" is bigger than array length return null.
	if n > len(x) {
		return 0
	}

	// Create variable "currVal" and save the sum
	// Create variable "tempVal" and compare with currVal
	var currVal int
	var tempVal int

	// Take the number of elements "n" from the array and sum them.
	for _, v := range x[0:n] {
		currVal += v
	}

	tempVal = currVal

	// Create a loop starting at "n" position in the array
	for i := n; i < len(x); i++ {
		tempVal = tempVal - x[i-n] + x[i]
		if tempVal > currVal {
			currVal = tempVal
		}
	}

	// 	Check for the maxVal compared with currVal
	return currVal
}
