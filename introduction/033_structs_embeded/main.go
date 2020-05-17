/* =====================================================================
 * -- Name ------ : Structs embeded
 * -- Date ------ : May 16, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : Practicing structs by creating onw and using it in another.
 ===================================================================== */

package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	typeName  vehicle
	fourWheel bool
}

type sedan struct {
	typeName vehicle
	luxury   bool
}

func main() {
	vs := sedan{
		typeName: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}

	vt := truck{
		typeName: vehicle{
			doors: 2,
			color: "grey",
		},
		fourWheel: false,
	}

	fmt.Println(vs)
	fmt.Println(vs.typeName.doors)

	fmt.Println(vt)
	fmt.Println(vt.typeName.doors)
}
