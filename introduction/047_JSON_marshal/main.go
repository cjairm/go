/* =====================================================================
 * -- Name ------ : JSON marshal
 * -- Date ------ : Jun 12, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : marshal the []user to JSON
 ===================================================================== */

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "Tonny",
		Age:   32,
	}

	u2 := user{
		First: "Barry",
		Age:   27,
	}

	u3 := user{
		First: "Bruce",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	usersMarshal, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(usersMarshal))
}
