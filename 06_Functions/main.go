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
// HOF - High Ordered function
// Functions which
// a) function has function as input parameter
// b) function which returns another function

// ********************************************

// Accept function as input parameter
func print(x, y int, area func(int, int) int) {
	fmt.Printf("Area is: %d\n", area(x, y))
}

// Returns a function
func getAreaFunc() func(int, int) int {
	return func(x, y int) int {
		return x * y
	}
}

// Others
func testfunction6(value string, callback func(string)) {
	value += " - 42"
	callback(value)
}

// ********************************************
// INNER FUNCTION CALL
// ********************************************
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
