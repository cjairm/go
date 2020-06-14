/* =====================================================================
 * -- Name ------ : Web Test
 * -- Date ------ : Janu 14, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description : First testing for web development
 ===================================================================== */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
