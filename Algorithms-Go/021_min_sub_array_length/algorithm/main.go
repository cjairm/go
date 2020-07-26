// Package algorithm contains different algorithms
package algorithm

// MinSubArrayLength search for the max length that sum a num
func MinSubArrayLength(xi []int, n int) int {
	if len(xi) <= 1 {
		return xi[len(xi)-1]
	}
	// Init helper variables
	var eleCount, curTotal, i, j int
	// Init variable that tends to infinity
	minLength := 999999999999999999

	// Init loop while first pointer is lower than length of the array
	for len(xi) > i {
		// 	If current total sum is lower than the number expected

		if curTotal < n {
			// Check that length of the second pointer and break if necessary
			if j > len(xi)-1 {
				break
			}
			// Increment current sum by adding the current element (array[position])
			curTotal += xi[j]
			// Increment the count of the elements
			eleCount++
			// Increment the second pointer
			j++
		} else {
			// 	If the current sum is equal or higher than the number expected
			if eleCount < minLength {
				minLength = eleCount
			}
			// If the count of elements is lower than the variable that tends to infinity, save that count in the infinity variable
			eleCount = 0
			// Reset count of elements, current total and increment the first pointer
			curTotal = 0
			// Then make equal the second pointer to the first one
			i++
			j = i
		}
	}

	// Just in case the infinity variable never change, return zero
	if minLength == 999999999999999999 {
		return 0
	}

	// Return infinity variable
	return minLength
}
