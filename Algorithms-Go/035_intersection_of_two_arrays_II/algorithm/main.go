// Package algorithm contains different algorithms
package algorithm

import (
	"math"
)

// Intersect Given two arrays, write a function to compute their intersection.
func Intersect(xi1, xi2 []int) []int {
	if len(xi1) <= 0 {
		return []int{}
	} else if len(xi2) <= 0 {
		return []int{}
	}

	var i, j int
	var result = []int{}
	xi1 = MergeSort(xi1)
	xi2 = MergeSort(xi2)

	if len(xi2) > len(xi1) {
		tmp := xi1
		xi1 = xi2
		xi2 = tmp
	}

	for len(xi1) > i && len(xi2) > j {
		numberOne := xi1[i]
		numberTwo := xi2[j]
		if numberOne == numberTwo {
			result = append(result, numberOne)
			j++
			i++
		} else if xi2[j] > xi1[i] {
			i++
		} else {
			j++
		}
	}

	return result
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
