// Package algorithm contains different algorithms
package algorithm

import (
	"math"
)

// MergeSort sorts an array of numbers
// Create MergeSort function
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
	// Create an empty array, take a look at the smallest values in each input array
	xxi := []int{}
	var i, j int

	// While there are still values we haven't looked at...
	for i < len(xi1) && j < len(xi2) {
		if xi2[j] > xi1[i] {
			// If the value in the first array is smaller than the value in the second array,
			// push the value in the first array into our results and move on to the next value in the first array
			xxi = append(xxi, xi1[i])
			i++
		} else {
			// If the value in the first array is larger than the value in the second array,
			// push the value in the second array into our results and move on to the next value in the second array
			xxi = append(xxi, xi2[j])
			j++
		}
	}

	// Once we exhaust one array, push in all remaining values from the other array
	for i < len(xi1) {
		xxi = append(xxi, xi1[i])
		i++
	}

	for j < len(xi2) {
		xxi = append(xxi, xi2[j])
		j++
	}

	return xxi
}
