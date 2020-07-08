// Package algorithm contains different algorithms
package algorithm

// SameButSquared compare two slices and creck if they are the same but squared
// Create a function called "SameButSquared"
func SameButSquared(x1 []int, x2 []int) bool {
	// If the length are different stop the script and return false
	if len(x1) != len(x2) {
		return false
	}

	// Create frequency variable helper called fq1
	fq1 := make(map[int]int)

	// Create another frequency variable helper called fq2
	fq2 := make(map[int]int)

	// Iterate first array
	for _, v1 := range x1 { // O(n)
		// If element of the array exists in fq1

		// Increment the key of fq1

		// If element of the array NOT exists in fq1

		// Create new element in fq1 with key name same as the element of the array
		fq1[v1]++
	}

	// Iterate second array
	for _, v2 := range x2 { // O(n)
		// If element of the array exists in fq2

		// Increment the key of fq2

		// If element of the array NOT exists in fq2

		// Create new element in fq2 with key name same as the element of the array
		fq2[v2]++
	}

	// Iterate fq1 helper that was created
	for k := range fq1 { // O(n)
		// If the current key value does NOT exists the array fq2 stop the script and return false
		if fq2[k*k] == 0 {
			return false
		}

		// If the current value does NOT match with same value but squared stop the script and retur false
		if fq1[k] != fq2[k*k] {
			return false
		}
	}

	return true
}
