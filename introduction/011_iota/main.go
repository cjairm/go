/* =====================================================================
 * -- Name ------ : IOTA
 * -- Date ------ : May 8, 2020
 * -- Author ---- : Carlos Mendez
 ===================================================================== */

package main

import "fmt"

const (
	g = 2020
	h = 2020 + iota
	i = 2020 + iota
	j = 2020 + iota
)

func main() {
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
}
