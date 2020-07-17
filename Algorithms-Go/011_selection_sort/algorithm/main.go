// Package algorithm contains different algorithms
package algorithm

// SelectionSort sorts an array of numbers
// Create SelectionSort function
func SelectionSort(xi []int) []int {
	for i := 0; i < len(xi); i++ { // O(n)
		// Store the first element as the smallest value you've seen so far.
		lowest := i
		for j := i + 1; j < len(xi); j++ { // O(n)
			// Compare this item to the next item in the array until you find a smaller number.
			if xi[j] < xi[lowest] {
				// If a smaller number is found, designate that smaller number to be the new "minimum" and continue until the end of the array.
				lowest = j
			}
		}

		// If the "minimum" is not the value (index) you initially began with, swap the two values.
		if i != lowest {
			//SWAP!
			xi[i], xi[lowest] = xi[lowest], xi[i]
		}
		// Repeat this with the next element until the array is sorted.
	}

	return xi // O(n^2)
}
