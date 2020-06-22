/* =====================================================================
 * -- Name ------ : Anonymous Structs
 * -- Date ------ : May 16, 2020
 * -- Author ---- : Carlos Mendez
 ===================================================================== */

package main

import "fmt"

func main() {

	myStruct := struct {
		superHero  string
		copycats   map[string]string
		superPower []string
	}{
		superHero: "Batman",
		copycats: map[string]string{
			"Florida": "Iron Man",
			"Stark":   "Arrow",
		},
		superPower: []string{
			"Rich",
			"Humanity",
		},
	}

	fmt.Println(myStruct.superPower)
	fmt.Println(myStruct.copycats)
	fmt.Println(myStruct.superPower)

	for k, v := range myStruct.copycats {
		fmt.Println(k, v)
	}

	for i, val := range myStruct.superPower {
		fmt.Println(i, val)
	}
}
