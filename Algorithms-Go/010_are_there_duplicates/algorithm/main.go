// Package algorithm contains different algorithms
package algorithm

// AreThereDuplicates search for duplicate values
// Create AreThereDuplicates function
func AreThereDuplicates(xi ...int) bool {
	// Sort xi
	xi = BubbleSort(xi)

	// Init two helpers (pointers) h1, h2
	var h1 int
	h2 := 1

	// Loop xi until the pointer is greater than the length of xi
	for h2 < len(xi) {
		// If the value at h1 it's equal to h2
		if xi[h1] == xi[h2] {
			return true
		}
		h1++
		h2++
	}

	return false
}

// BubbleSort sorts an array of numbers
func BubbleSort(xi []int) []int {
	var flag bool
	for i := len(xi); i > 0; i-- {
		flag = true
		for j := 0; j < i-1; j++ {
			if xi[j] > xi[j+1] {
				xi[j], xi[j+1] = xi[j+1], xi[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}

	return xi
}
