// Package algorithm contains different algorithms
package algorithm

import "math"

// SearchNumber search for a number in an array
// Create a function called "SearchNumber"
func SearchNumber(x []int, n int) int {

	// If the array received contains zero elements stop the program ans return -1
	// If the array received contains one element stop the program and return 1
	if len(x) == 1 {
		return 1
	} else if len(x) == 0 {
		return -1
	}

	// Init a pointer at the end of the array and another one at the beggining of the array.
	var min int
	max := len(x) - 1

	// Init a loop while max pointer is bigger than min pointer
	for min <= max {
		// Get the middle of the array
		mid := int(math.Floor(float64(min)+float64(max)) / 2)
		currEle := x[mid]

		// If middle value is bigger than searched number move max pointer to left (middle pointer - 1)
		// If middle value is lower than searched number move min pointer to right (middle pointer + 1)
		// if Equal return position (middle index)
		if currEle < n {
			min = mid + 1
		} else if currEle > n {
			max = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
