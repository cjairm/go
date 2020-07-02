package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isAnagramValid("Affbb", "abbfF"))
}

// create function called "isAnagramValid"
func isAnagramValid(s1 string, s2 string) bool {
	// if both strings has the same length. We continue
	if len(s1) != len(s2) {
		return false
	}
	// if both strings are the same. We stop.
	if s1 == s2 {
		return true
	}
	// declare helpers to save each variable
	h1 := make(map[string]int)
	h2 := make(map[string]int)

	// Iterate first string and save how many letters exists
	// First we convert the string one to lowercase
	for _, currLetter := range strings.ToLower(s1) {
		// fmt.Printf("Start Index: %d Value:%s\n", i, string(currLetter))
		// It start always at zero. Zero Values
		h1[string(currLetter)]++
	}

	// Iterate second string and save how many letter exists
	// First we convert the string two to lowercase
	for _, currLetter := range strings.ToLower(s2) {
		// fmt.Printf("Start Index: %d Value:%s\n", i, string(currLetter))
		// It start always at zero. Zero Values
		h2[string(currLetter)]++
	}

	// Iterate one of the object created.
	for k := range h1 {
		// if the letter exists. Continue
		if h1[k] == 0 || h2[k] == 0 {
			return false
		}

		// If counter of characters are the same. Continue
		if h1[k] != h2[k] {
			return false
		}
	}

	// Return true
	return true
}
