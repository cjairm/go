// Package algorithm contains different algorithms
package algorithm

import "fmt"

// RemoveConsecutiveChar remove all consecutive chars
func RemoveConsecutiveChar(str string, k int) string {
	s := 1
	f := 0
	count := 0
	var result string

	for len(str) > f {
		fmt.Println(f, s, str)
		if string(str[f]) != string(str[s]) {
			f++
			count = 0
		} else {
			count++
			s++
			if s >= len(str)-1 {
				f++
				s = f + 1
				fmt.Println("===", f, len(str), count)
				if f >= len(str)-1 {
					if count == k {
						result = str[:f] + str[s:len(str)]
						return RemoveConsecutiveChar(result, k)
					} else {
						break
					}
				}
			}
		}

		if count == k {
			result = str[:f] + str[s:len(str)]
			return RemoveConsecutiveChar(result, k)
		}
	}
	return str
}
