// Package algorithm contains different algorithms
package algorithm

// Power accepts a base and an exponent
func Power(b, e int) int {
	if e == 0 {
		return 1
	}
	return b * Power(b, e-1)
}
