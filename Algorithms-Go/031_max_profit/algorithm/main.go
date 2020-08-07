// Package algorithm contains different algorithms
package algorithm

// MaxProfit finds the maximum profit.
func MaxProfit(xi []int) int {
	var have bool
	var tomorrow, today, profit int
	for i := 0; i < len(xi); i++ {
		today = xi[i]
		if i+1 >= len(xi) {
			tomorrow = 0
		} else {
			tomorrow = xi[i+1]
		}

		if have {
			if today > tomorrow {
				profit += today
				have = false
			}
		} else {
			if tomorrow > today {
				have = true
				profit -= today
			}
		}
	}
	if have {
		return profit + today
	}
	return profit

}
