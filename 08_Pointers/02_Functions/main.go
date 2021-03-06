package main

import (
	"fmt"
)

func main() {

	proofChangeMe2()

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
// ********************************************
// VS
// ********************************************
// ********************************************

// ********************************************
// Passing reference - value
// ********************************************

// !!! int,string,rune ... etc. DON'T CHANGED !!!!

func changeMe2(z *string) {
	fmt.Println(z) // Todd
	*z = "Rocky"
	fmt.Println(z) // Rocky
}

func proofChangeMe2() {
	testString := "Adam"
	changeMe1(testString)
	fmt.Println(testString) // Adam
}

// ********************************************
// Passing reference - struct
// ********************************************

// !!! struct ... IS CHANGED !!!!

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

	fmt.Println(c1) // &{Rocky 44}
}
