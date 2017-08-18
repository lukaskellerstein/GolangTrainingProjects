package main

import (
	"fmt"
	"math"
)

// example of empty interface
type ancestor interface{}

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

func (c circle) calculateArea() float64 {
	return math.Pi * (float64(c.radius) * float64(c.radius))
}

// ****************
// STRUCT circle implements shape interface
// ****************
type square struct {
	side int
}

func (c square) calculateArea() float64 {
	return float64(c.side) * float64(c.side)
}

//METHOD FOR ANY SHAPE ------------------------
func totalArea(shapes ...shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.calculateArea()
	}
	return area
}

//METHOD FOR ANYONE - Empty interface ------------------------
func printMyType(anyone ancestor) {
	fmt.Printf("%T \n", anyone)
}

// ****************
// ****************
// MAIN
// ****************
// ****************
func main() {
	circle1 := circle{2}
	fmt.Println(circle1.calculateArea())

	square1 := square{10}
	fmt.Println(square1.calculateArea())

	//OR

	fmt.Println(totalArea(circle1))
	fmt.Println(totalArea(square1))
}
