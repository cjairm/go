// Package algorithm contains different algorithms
package algorithm

// PlusOne increment one to the integer.
func PlusOne(xi []int) []int {
	digLen := len(xi)
	var carry int
	currDigit := xi[digLen-1]
	if currDigit < 9 {
		xi[digLen-1]++
	} else {
		carry++
		xi[digLen-1] = 0
	}

	for i := digLen - 2; i >= 0; i-- {
		if carry > 0 {
			if xi[i]+carry <= 9 {
				xi[i] = xi[i] + carry
				carry = 0
			} else {
				carry = 1
				xi[i] = 0
			}
		}
	}

	if carry > 0 {
		xxi := []int{}
		xxi = append(xxi, carry)
		xxi = append(xxi, xi...)
		return xxi
	}

	return xi
}
