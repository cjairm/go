// Package algorithm contains different algorithms
package algorithm

// MaxSumSubArray search for the max sum in a array
func MaxSumSubArray(xi []int, n int) int {
	if n > len(xi) {
		return 0
	}

	// Save in a variable called currentSum the sum of the first n elements of the array
	var currentSum, tempSum int
	i := 1
	for _, v := range xi[:n] {
		currentSum = currentSum + v
	}

	// Create a loop while pointer is lower than the length of the array
	for len(xi)-n+1 > i {
		// fmt.Println(i)
		// Save in temporal variable the sum of the array from pointer to pointer plus n
		tempSum = 0
		for _, curVal := range xi[i : i+n] {
			tempSum = tempSum + curVal
		}

		// If temporal sum is bigger than currentSum make it equal to temporal sum
		if tempSum > currentSum {
			currentSum = tempSum
		}

		// Increment pointer
		i++
	}

	return currentSum
}
