// Package algorithm contains different algorithms
package algorithm

// QuickSort as the name says sort the array
func QuickSort(xi []int) []int {
	if len(xi) < 2 {
		return xi
	}

	left, right := 0, len(xi)-1

	var pivot int

	xi[pivot], xi[right] = xi[right], xi[pivot]

	for i := range xi {
		if xi[i] < xi[right] {
			xi[left], xi[i] = xi[i], xi[left]
			left++
		}
	}

	xi[left], xi[right] = xi[right], xi[left]

	QuickSort(xi[:left])
	QuickSort(xi[left+1:])

	return xi
}
