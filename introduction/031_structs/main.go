/* =====================================================================
 * -- Name ------ : Structs
 * -- Date ------ : May 16, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : Creating first struct.
 ===================================================================== */

package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	iceCreamFlavors []string
}

func main() {
	people := []person{{
		firstName: "Carlos",
		lastName:  "Mendez",
		iceCreamFlavors: []string{
			"Vainilla",
			"Chocolate",
		},
	}, {
		firstName: "Jair",
		lastName:  "Juarez",
		iceCreamFlavors: []string{
			"Coco",
			"Pistache",
		},
	}}

	for _, personValue := range people {
		fmt.Printf("Hi, my name is %v and my ice cream favorites flavors are: \n", personValue.firstName+" "+personValue.lastName)
		for _, iceCreams := range personValue.iceCreamFlavors {
			fmt.Println(iceCreams)
		}
	}
}
