/* =====================================================================
 * -- Name ------ : switch w/statement
 * -- Date ------ : May 8, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : working for first time with switch case and statement
 ===================================================================== */

package main

import "fmt"

func main() {
	myFavSport := "soccer"
	switch myFavSport {
	case "Bei":
		fmt.Println("this is NOT printed")
	case "Bas":
		fmt.Println("this is NOT printed")
	case "soccer":
		fmt.Printf("%v is my favorite sport \n", myFavSport)
	}
}
