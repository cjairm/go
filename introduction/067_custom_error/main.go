package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	v, err := sqrt(-10.23)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(v)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		err := fmt.Errorf("The sqrt value has to be greater than 0")
		return 0, sqrtError{
			"50.2289 N",
			"99.4656 W",
			err,
		}
	}
	return 42, nil
}
