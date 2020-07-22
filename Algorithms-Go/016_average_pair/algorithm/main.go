// Package algorithm contains different algorithms
package algorithm

import (
	"math"
)

// AveragePair search for a pair of numbers that return specific number
func AveragePair(xi []int, n float64) bool {
	// Sort the array
	xi = MergeSort(xi)

	// Create two helper pointers
	f := 0
	s := len(xi) - 1
	var avrg float64
	// Start looping the array while the end pointer is bigger than start pointer
	for s > f {
		// 	Calculate the average
		avrg = float64((xi[f] + xi[s])) / 2.0

		if avrg == n {
			// 	If average is equal to the num searched return true
			return true
		} else if n > avrg {
			// 	If num is bigger than average, so move start pointer to the rigth
			f++
		} else {
			// 	If average is bigger than num, so move end pointer to the left
			s--
		}
	}

	return false
}

// MergeSort sorts an array of numbers
func MergeSort(xi []int) []int {
	if len(xi) <= 1 {
		return xi
	}
	mid := int(math.Floor(float64(len(xi) / 2)))
	left := MergeSort(xi[0:mid])
	right := MergeSort(xi[mid:])
	return MergeSortedArrays(left, right)
}

// MergeSortedArrays as the name says merge two arrays
func MergeSortedArrays(xi1, xi2 []int) []int {
	xxi := []int{}
	var j, i int

	for i < len(xi1) && j < len(xi2) {
		if xi2[j] > xi1[i] {
			xxi = append(xxi, xi1[i])
			i++
		} else {
			xxi = append(xxi, xi2[j])
			j++
		}
	}

	if len(xi1) > i {
		xxi = append(xxi, xi1[i:]...)
	}

	if len(xi2) > j {
		xxi = append(xxi, xi2[j:]...)
	}

	return xxi
}
