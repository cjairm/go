// Package algorithm contains different algorithms
package algorithm

// GetPivotQuickSort as the name says merge two arrays
// Create GetPivotQuickSort function
func GetPivotQuickSort(xi []int, start int, end int) (int, []int) {
	if len(xi) <= 1 {
		return len(xi), xi
	}

	// It will help to accept three arguments: an array, a start index,
	// and an end index (these can default to 0 and the array length minus 1, respectively)
	if end == 0 {
		end = len(xi) - 1
	}

	// Grab the pivot from the start of the array
	pivot := xi[start]

	// Store the current pivot index in a variable (this will keep track of where the pivot should end up)
	swapIndex := start

	// Loop through the array from the start until the end
	for i := start + 1; i <= end; i++ {
		// If the pivot is greater than the current element, increment the pivot index variable
		// and then swap the current element with the element at the pivot index
		if pivot > xi[i] {
			swapIndex++
			xi[i], xi[swapIndex] = xi[swapIndex], xi[i]
		}
	}

	// Swap the starting element (i.e. the pivot) with the pivot index
	xi[start], xi[swapIndex] = xi[swapIndex], xi[start]

	// Return the pivot index
	return swapIndex, xi
}
