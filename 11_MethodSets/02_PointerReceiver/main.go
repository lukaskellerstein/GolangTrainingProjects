package main

import (
	"fmt"
	"math"
)

// ****************
// INTERFACE shape
// ****************
type shape interface {
	calculateArea() float64
}

// ****************
// STRUCT circle implements shape interface
// ****************
type circle struct {
	radius int
}

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// POINTER RECEIVER
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func (c *circle) calculateArea() float64 {
	return math.Pi * (float64(c.radius) * float64(c.radius))
}

// helper method
func print(s shape) {
	fmt.Println("area", s.calculateArea())
}

// ****************
// ****************
// MAIN
// ****************
// ****************
func main() {
	c := circle{5}
	//pointer type
	print(&c)
	//value type - CANNOT BE USED !!!!!!
	//print(c)
}
