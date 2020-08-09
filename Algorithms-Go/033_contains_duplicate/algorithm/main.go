// Package algorithm contains different algorithms
package algorithm

import (
	"math"
)

// ContainsDuplicate find if the array contains any duplicates
func ContainsDuplicate(xi []int) bool {
	var i, j int
	j = 1
	xxi := MergeSort(xi)

	for len(xxi) > j {
		if xxi[i] == xxi[j] {
			return true
		}
		i++
		j++
	}

	return false
}

// ContainsDuplicateTwo find if the array contains any duplicates
func ContainsDuplicateTwo(xi []int) bool {
	var i, j int
	j = 1

	for len(xi) > i {
		if j >= len(xi) {
			i++
			j = i + 1
			if j >= len(xi) {
				break
			}
		}

		if xi[i] == xi[j] {
			return true
		}
		j++
	}

	return false
}

// MergeSort sorts an array
func MergeSort(xi []int) []int {
	if len(xi) <= 1 {
		return xi
	}

	mid := int(math.Floor(float64(len(xi) / 2)))
	left := MergeSort(xi[0:mid])
	right := MergeSort(xi[mid:])

	return MergeArrays(left, right)
}

// MergeArrays as the name says merge two arrays
func MergeArrays(xi1, xi2 []int) []int {
	var xxi []int
	var i, j int

	for len(xi1) > i && len(xi2) > j {
		if xi2[j] > xi1[i] {
			xxi = append(xxi, xi1[i])
			i++
		} else {
			xxi = append(xxi, xi2[j])
			j++
		}
	}

	for len(xi1) > i {
		xxi = append(xxi, xi1[i])
		i++
	}

	for len(xi2) > j {
		xxi = append(xxi, xi2[j])
		j++
	}

	return xxi
}
