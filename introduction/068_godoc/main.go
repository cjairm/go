package main

import (
	"fmt"

	"github.com/cjairm/introduction/068_godoc/dog"
)

type d struct {
	name  string
	years int
}

func main() {
	k := d{
		name:  "Killer",
		years: dog.Years(8),
	}

	fmt.Println(k)
}
