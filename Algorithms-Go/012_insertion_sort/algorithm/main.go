// Package algorithm contains different algorithms
package algorithm

// InsertionSort sorts an array of numbers
// Create InsertionSort function
func InsertionSort(xi []int) []int {
	var currVal int
	for i := 1; i < len(xi); i++ { // O(n)
		// Start by picking the second element in the array
		currVal = xi[i]
		// Now compare the second element with the one before it and swap if necessary.
		j := i - 1
		for j > -1 && xi[j] > currVal { // O(n)
			// Continue to the next element and if it is in the incorrect order,
			// iterate through the sorted portion (i.e. the left side) to place the element in the correct place.
			xi[j+1] = xi[j]
			j--
		}
		xi[j+1] = currVal
		// Repeat until the array is sorted.
	}

	return xi // O(n^2)
}
