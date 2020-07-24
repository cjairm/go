// Package algorithm contains different algorithms
package algorithm

// QuickSort as the name says sort the array
func QuickSort(xi []int, left int, right int) []int {
	if right == 0 {
		right = len(xi) - 1
	}

	if left < right {
		pivotIndex, _ := GetPivotQuickSort(xi, left, right)

		QuickSort(xi, left, pivotIndex-1)

		QuickSort(xi, pivotIndex+1, right)
	}
	
	return xi
}

// GetPivotQuickSort as the name says merge two arrays
func GetPivotQuickSort(xi []int, start int, end int) (int, []int) {
	if len(xi) <= 1 {
		return len(xi), xi
	}

	if end == 0 {
		end = len(xi) - 1
	}

	pivot := xi[start]
	swapIndex := start

	for i := start + 1; i <= end; i++ {
		if pivot > xi[i] {
			swapIndex++
			xi[i], xi[swapIndex] = xi[swapIndex], xi[i]
		}
	}

	xi[start], xi[swapIndex] = xi[swapIndex], xi[start]

	return swapIndex, xi
}
