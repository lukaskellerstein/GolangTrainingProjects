package main

import (
	"fmt"
)

func main() {

}

// ********************************************
// Basics - accept and return values
// ********************************************

// Two params function
func testfunction1(value1 string, value2 string) {
	fmt.Println(value1 + " - " + value2)
}

// Two params function with return value
func testfunction2(value1 string, value2 string) string {
	return value1 + " - " + value2
}

// Two params function with return two value
func testfunction3(value1 string, value2 string) (string, string) {
	return value1, value2
}

// N params function with return array of values
func testfunction4(values ...string) []string {
	return values
}

// array params function with return array of values
func testfunction5(values []string) []string {
	return values
}

// ********************************************
// Callbacks - accept and return functions
// ********************************************

func testfunction6(value string, callback func(string)) {
	value += " - 42"
	callback(value)
}

func proofTestfunction6() {
	testfunction6("Number of whole universe", func(text string) {
		fmt.Println(text)
	})
}

// ********************************************
// Passing value
// ********************************************

func changeMe1(z string) {
	fmt.Println(z) // Todd
	z = "Rocky"
	fmt.Println(z) // Rocky
}

func proofChangeMe1() {
	testString := "Adam"
	changeMe1(testString)
}

// ********************************************
// Passing reference - value
// ********************************************

func changeMe2(z *string) {
	fmt.Println(z) // Todd
	*z = "Rocky"
	fmt.Println(z) // Rocky
}

func proofChangeMe2() {
	testString := "Adam"
	changeMe1(testString)
}

// ********************************************
// Passing reference - struct
// ********************************************

type customer struct {
	name string
	age  int
}

func changeMe3(z *customer) {
	fmt.Println(z)       // &{Todd 44}
	fmt.Println(&z.name) // 0x8201e4120
	z.name = "Rocky"
	fmt.Println(z)       // &{Rocky 44}
	fmt.Println(&z.name) // 0x8201e4120
}

func proofChangeMe3() {
	c1 := customer{"Todd", 44}

	changeMe3(&c1)
}

// ********************************************
// Defer
// ********************************************

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func proofDefer() {

	world()
	hello()

	defer world()
	hello()

}
