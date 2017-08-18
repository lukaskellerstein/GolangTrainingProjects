package main

import (
	"fmt"
)

var a string
var b int
var c, d, e int
var f float32
var g bool

var za, zb, zc int = 1, 2, 3
var ya, yb, yc = 1, 2, 3

var aa = 5

func main() {
	//DECLARE

	//ASSIGN
	a = "someString"
	b = 112313213
	c, d, e = 1, 2, 3

	//INITIALIZATION

	ab := 5

	za = 1
	zb = 2
	zc = 3

	ya = 1
	yb = 2
	yc = 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(ab)
	fmt.Println(za)
	fmt.Println(zb)
	fmt.Println(zc)
	fmt.Println(ya)
	fmt.Println(yb)
	fmt.Println(yc)
}
