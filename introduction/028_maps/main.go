/* =====================================================================
 * -- Name ------ : Maps
 * -- Date ------ : May 10, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : Create maps - This means arrays with key as selectors
 ===================================================================== */

package main

import "fmt"

func main() {
	////myMap := map[type]TYPE_CONTENT{content}
	myMap := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	fmt.Println(myMap)

	for keyName, mapValue := range myMap {
		fmt.Println(keyName)
		for _, ele := range mapValue {
			fmt.Printf("\t %v \n", ele)
		}
	}
}
