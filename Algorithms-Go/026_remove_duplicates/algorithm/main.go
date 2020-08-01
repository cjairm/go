// Package algorithm contains different algorithms
package algorithm

import (
	"math"
)

// RemoveDuplicates remove the duplicates in-place such that each element appear only once
func RemoveDuplicates(xi []int) ([]int, int) {
	if len(xi) == 0 {
		return []int{}, 0
	} else if len(xi) == 1 {
		return xi, 1
	} else if len(xi) == 2 {
		if xi[0] == xi[1] {
			return []int{xi[0]}, 1
		}
		return xi, 2

	}

	var xxi []int
	// Sort array
	xxi = Sort(xi)

	// Init helper variables (two pointers)
	i := 0
	j := 1

	// Start loop until first pointer reach the end of the array
	for i < len(xxi) && j < len(xxi) {
		// If first pointer is equal to the second pointer move pointer two until find diferent character
		if xxi[i] == xxi[j] {
			j++
		} else {
			// Otherwise move two pointers and swap them.
			i++
			xxi[i], xxi[j] = xxi[j], xxi[i]
			j++
		}
	}

	index := i + 1

	return xxi[0:index], len(xxi[0:index])
}

// Sort sorts an array of numbers
func Sort(xi []int) []int {
	if len(xi) <= 1 {
		return xi
	}

	mid := int(math.Floor(float64(len(xi) / 2.0)))
	left := Sort(xi[0:mid])
	right := Sort(xi[mid:])
	return MergeSlices(left, right)
}

// MergeSlices as the name says merge two arrays
func MergeSlices(xi1, xi2 []int) []int {
	var xxi []int
	var i, j int

	for i < len(xi1) && j < len(xi2) {
		if xi1[i] < xi2[j] {
			xxi = append(xxi, xi1[i])
			i++
		} else {
			xxi = append(xxi, xi2[j])
			j++
		}
	}

	xxi = append(xxi, xi1[i:]...)
	xxi = append(xxi, xi2[j:]...)

	return xxi
}
