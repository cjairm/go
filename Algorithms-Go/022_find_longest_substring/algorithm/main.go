// Package algorithm contains different algorithms
package algorithm

// FindLongestSubstring returns the length of the longest substring with all distinct characters.
func FindLongestSubstring(s string) int {
	letterSeen := make(map[string]int)
	var i, j, currTotal, maxLength int

	for len(s) > i {
		if j >= len(s) {
			i++
			j = i
			currTotal = 0
		} else {
			currLetter := string(s[j])
			if letterSeen[currLetter] == 0 {
				letterSeen[currLetter]++
				currTotal++
				j++
			} else {
				i++
				j = i
				currTotal = 0
			}

			if currTotal > maxLength {
				maxLength = currTotal
			}
		}
		i++
	}

	return maxLength
}
