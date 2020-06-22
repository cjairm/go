/* =====================================================================
 * -- Name ------ : Structs and map
 * -- Date ------ : May 16, 2020
 * -- Author ---- : Carlos Mendez
 ===================================================================== */

package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	iceCreamFlavors []string
}

func main() {
	p1 := person{
		firstName: "Carlos",
		lastName:  "Mendez",
		iceCreamFlavors: []string{
			"Vainilla",
			"Chocolate",
		},
	}

	p2 := person{
		firstName: "Jair",
		lastName:  "Juarez",
		iceCreamFlavors: []string{
			"Coco",
			"Pistache",
		},
	}

	myMap := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, personValue := range myMap {
		fmt.Printf("Hi, my name is %v and my ice cream favorites flavors are: \n", personValue.firstName+" "+personValue.lastName)
		for _, iceCreams := range personValue.iceCreamFlavors {
			fmt.Println(iceCreams)
		}
	}
}
