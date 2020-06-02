/* =====================================================================
 * -- Name ------ : Functions attach excercise
 * -- Date ------ : Jun 2, 2020
 * -- Author ---- : Carlos Mendez
 * -- Description :
	● create a type ​ SQUARE
	● create a type ​ CIRCLE
	● attach a method to each that calculates ​ AREA​ and returns it
		○ circle area= π r 2
		○ square area = L * W
	● create a type ​ SHAPE​ that defines an interface as anything that has the ​ AREA​ method
	● create a func ​ INFO​ which takes type shape and then prints the area
	● create a value of type square
	● create a value of type circle
	● use func info to print the area of square
	● use func info to print the area of circle
 ===================================================================== */

package main

import (
	"fmt"
	"math"
	"strconv"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

type shape interface {
	area() float64
}

func info(s shape) string {
	return "My area is " + strconv.FormatFloat(s.area(), 'f', 6, 64)
}

func main() {
	sq := square{
		length: 12.10,
	}

	ci := circle{
		radius: 3,
	}

	fmt.Println(info(sq))
	fmt.Println(info(ci))
}
