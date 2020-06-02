/* =====================================================================
 * -- Name ------ : Anonymous functions
 * -- Date ------ : Jun 2, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : Build and use an anonymous func
 ===================================================================== */

package main

import "fmt"

func main() {
	func(n string) {
		fmt.Println("Hello my name is", n)
	}("Carlos Mendez")
}
