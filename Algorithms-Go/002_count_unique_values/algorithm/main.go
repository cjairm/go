// Package algorithm contains different algorithms 
package algorithm

// GetCountOfUniqueValues take a slice of ints and check for unique values
// Create a function called "getCountOfUniqueValues"
func GetCountOfUniqueValues(x []int) int {
	// Check the length of the array

	// If array length is equal one/zero, return the count of one/zero unique value
	if len(x) <= 1 {
		return len(x)
	}

	// If array length is equal two, check the first and second value. Then, return the count of unique values
	if len(x) == 2 {
		if x[0] == x[1] {
			return 1
		}
		return 2
	}

	// Create reference (pointer) at the begining (position 1) of the array called firstRef
	var firstRef int

	// Create another reference (pointer) at position 2 in the array called secondRef
	secondRef := 1

	// Create loop. Stay there until reference two is equal to the length of the array
	for {
		// If the element in the array that points firstRef and that one that points secondref two are equal. Incremente secondRef
		// If are diferents increment secondRef and firstRef
		if x[firstRef] != x[secondRef] {
			firstRef++
			x[firstRef] = x[secondRef]
		}
		secondRef++

		if secondRef >= len(x) {
			break
		}
	}

	// Return value of firstRef
	return firstRef + 1
}
