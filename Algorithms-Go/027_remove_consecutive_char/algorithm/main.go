// Package algorithm contains different algorithms
package algorithm

// RemoveConsecutiveChar remove all consecutive chars
func RemoveConsecutiveChar(str string, k int) string {
	s := 1
	f := 0
	count := 0
	var result string

	for len(str) > s {
		if string(str[f]) != string(str[s]) {
			f++
			count = 0
		} else {
			count++
			s++
		}

		if count == k {
			result = str[:f] + str[s:len(str)]
			return RemoveConsecutiveChar(result, k)
		}
	}
	return str
}
