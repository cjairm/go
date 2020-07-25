// Package algorithm contains different algorithms
package algorithm

// IsSubsequence check if the first string has the same letter and same order as two.
func IsSubsequence(f, s string) bool {
	// Create helpers
	var j, i int

	if f == "" {
		return true
	}

	// Start loop while the second array pointer is bigger than length of array
	for len(s) > j {
		// 	If array1 at first pointer if equal to array2 at second pointer, increment the first pointer
		if f[i] == s[j] {
			i++
		}

		// 	If first pointer meet array1's length return true
		if i == len(f) {
			return true
		}

		// 	Increment second pointer
		j++
	}

	// Return false
	return false
}
