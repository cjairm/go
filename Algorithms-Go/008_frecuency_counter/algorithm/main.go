package algorithm

import (
	"strconv"
)

// SameFrecuency search for a frecuency in number
// Create SameFrecuency function
func SameFrecuency(n1, n2 int) bool {
	// Convert n1 and n2 to string
	s1 := strconv.Itoa(n1)
	s2 := strconv.Itoa(n2)

	if len(s1) != len(s2) {
		return false
	}

	// Create two helper slices xi1 to store n1 data and xi2 to store n2 data
	xi1 := make(map[string]int)
	xi2 := make(map[string]int)

	// Loop over n1 as string
	for _, v1 := range s1 {
		// Create key in the slice xi1 and / or increment value.
		currVal := string(v1)
		xi1[currVal]++
	}

	// Loop over n2 as string
	for _, v2 := range s2 {
		// Create key in the slice xi2 and / or increment value.
		currVal := string(v2)
		xi2[currVal]++
	}

	// Compare xi1 and xi2
	for k, v := range xi1 {
		if xi2[k] == 0 {
			return false
		}

		if xi2[k] != v {
			return false
		}
	}

	return true
}
