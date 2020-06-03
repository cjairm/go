/* =====================================================================
 * -- Name ------ : Callback function
 * -- Date ------ : Jun 3, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : A “callback” is when we pass a func into a func as an argument. For this exercise,
	● pass a func into a func as an argument
 ===================================================================== */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	g := greet(countWords, "Hello, I am Carlos Mendez")
	fmt.Println(g)
}

func countWords(s string) int {
	e := strings.Split(s, " ")
	return len(e)
}

func greet(count func(s string) int, s string) string {
	return `I said "` + s + `" with ` + strconv.Itoa(count(s)) + " words."
}
