// Package algorithm contains different algorithms
package algorithm

// Fib returns the number in the Fibonacci sequence
func Fib(i int) int {
	if i <= 2 {
		return 1
	}
	return Fib(i-1) + Fib(i-2)
}
