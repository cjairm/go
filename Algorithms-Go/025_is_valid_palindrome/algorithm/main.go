// Package algorithm contains different algorithms
package algorithm

import (
	"regexp"
	"strings"
)

// IsLetter checks if is alphanumeric
// var IsLetter = regexp.MustCompile(`/[^a-zA-Z0-9]/`).MatchString

// IsValidPalindrome determine if it is a palindrome
func IsValidPalindrome(s string) bool {
	// Check if string is empty return true
	if len(s) == 0 {
		return true
	}

	// Init helpers variables (pointers) to start from the beggining and the end of the array
	var left, right = 0, len(s) - 1

	// Start a loop while right pointer is bigger than left
	for right > left {
		currRight := strings.ToLower(string(s[right]))
		currLeft := strings.ToLower(string(s[left]))

		// if both pointers are alphanumeric
		if IsLetter(currRight) && IsLetter(currLeft) {
			// if pointers values are equal, increment left and decrement right pointers
			if currLeft == currRight {
				left++
				right--
			} else {
				return false
			}
		} else {
			// if left pointer is not alphanumeric increment this pointer
			if !IsLetter(currLeft) {
				left++
			}
			// if right pointer is not alphanumeric decrement this pointer
			if !IsLetter(currRight) {
				right--
			}
		}
	}

	return true
}

// IsLetter checks if is alphanumeric
func IsLetter(l string) bool {
	matched, _ := regexp.MatchString(`[^a-zA-Z0-9]`, l)
	return !matched
}
