package main

import (
	"fmt"
)

func main() {
	// ************************
	// Array
	// ************************
	var array1 [100]string // var - is used only for declaration -> need to be fill
	for i := 0; i < 100; i++ {
		array1[i] = string(i)
	}
	// ************************
	// Slice(= List) - created with initial values
	// ************************
	slice1 := []string{
		"string1",
		"string2",
		"string3",
	}
	for v := range slice1 {
		fmt.Printf("%v\n", v)
	}
	// Slice(= List) - created without initial values
	// length = capacity
	slice2 := make([]string, 3)
	// length and capacity
	// or slice3 := make([]string, 3, 3)
	// or slice4 := make([]string, 3, 100)
	// if i need add value into slice4 after index 3, i will need use append.
	// So I will have prepared only 3 piece of space in this slice

	// Slice - operations

	//append
	slice1 = append(slice1, "string4")
	slice2 = append(slice2, "string1")
	slice2 = append(slice2, "string2")
	slice2 = append(slice2, "string3")
	slice2 = append(slice2, "string4")

	//delete
	slice2 = remove(slice2, 3)

	//maps -----------------------------------------
	main2()

	fmt.Println("END")
}

func remove(slice []string, i int) []string {
	slice = append(slice[:i], slice[i+1:]...)
	return slice
}
