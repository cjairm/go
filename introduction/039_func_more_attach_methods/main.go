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
