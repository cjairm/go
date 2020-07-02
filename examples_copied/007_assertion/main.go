/* =====================================================================
 * -- Name ------ : Assertion
 * -- Date ------ : Jul 1, 2020
 * -- Author ---- : Todd McLeod
 * -- Description : Example
 * -- GitHub account ------ : @GoesToEleven
 ===================================================================== */

package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "need more coffee",
	}

	foo(c1)
}

func foo(e error) {
	// Example assertion: e.(customErr).info = we know that e has a type assigned
	fmt.Println("foo ran -", e, "\n", e.(customErr).info)
}
