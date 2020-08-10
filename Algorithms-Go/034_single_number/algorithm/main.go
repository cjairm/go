// Package algorithm contains different algorithms
package algorithm

import (
	"math"
)

// SingleNumber a non-empty array of integers, every element appears twice except for one. Find that single one.
func SingleNumber(xi []int) int {
	xxi := MergeSort(xi)
	var i int

	for len(xxi) > i {
		if i+1 >= len(xxi) {
			if xxi[i] != xxi[i-1] {
				return xxi[i]
			}
		}

		if xxi[i] == xxi[i+1] {
			i += 2
		} else {
			return xxi[i]
		}
	}

	return 0
}

// SingleNumberTwo a non-empty array of integers, every element appears twice except for one. Find that single one.
func SingleNumberTwo(xi []int) int {
	var i int
	mapCompare := make(map[int]int)

	for len(xi) > i {
		mapCompare[xi[i]]++
		i++
	}

	for k, v := range mapCompare {
		if v == 1 {
			return k
		}
	}

	return 0
}

// MergeSort sorts an array
func MergeSort(xi []int) []int {
	if len(xi) <= 1 {
		return xi
	}

	mid := int64(math.Floor(float64(len(xi) / 2)))
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
