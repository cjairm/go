// Package algorithm contains different algorithms
package algorithm

// SumZero look for two values in one slice than sum zero.
// Create a function called "SumZero"
func SumZero(x []int) []int {
	// Check the length of the array. If the length is zero or one return undefined
	if len(x) <= 0 {
		return nil
	}

	// Create one helper variable called hr that starts at the beggining of the array
	hl := 0

	// Create another helper variable called hl that starts at the end of the array
	hr := len(x) - 1

	// Create a loop while hr is greater than hl
	for hr > hl { // O(n)
		//     Sum two variables - array at position hr plus array at position hl
		sum := x[hr] + x[hl]

		if sum == 0 {
			//     If the sum is equal to zero return [array at position hr, array at position hl]
			return []int{x[hl], x[hr]}
		} else if sum > 0 {
			//     If the sum is greater than zero decrement in one hr
			hr--
		} else {
			//     If the sum is lower than zero increment in one hl
			hl++
		}
	}
	return nil
}
