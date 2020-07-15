package algorithm

// BubbleSort sorts an array of numbers
// Create BubbleSort function
func BubbleSort(xi []int) []int {
	var noSwapFlag bool
	// Start looping with a variable called i the end of the array towards the beggining
	for i := len(xi); i > 0; i-- { // O(n)
		noSwapFlag = true
		// Start an inner loop with a variable called j from the beggining until i - 1
		for j := 0; j < i-1; j++ { // O(n)
			// If arr[j] is greater|lower than arr[j+1], swap those values
			if xi[j] > xi[j+1] {
				// swap
				xi[j], xi[j+1] = xi[j+1], xi[j]
				noSwapFlag = false
			}
		}

		if noSwapFlag {
			break
		}
	}

	return xi // O(n^2)
}
