package main

import "fmt"

func main() {

	//variables --------------------------------------
	someTest := "There is some test text"

	fmt.Println("in variable someTest is: ", someTest)
	fmt.Println("and it is save in memory address: ", &someTest)

	// I don't know how to get back the string
	fmt.Printf("if i try translate it back to text: %d \n", &someTest)
	fmt.Printf("if i try translate it back to text: %s \n", &someTest)
	fmt.Printf("if i try translate it back to text: %p \n", &someTest)

	//array -----------------------------------------

	var array1 [10]string // var - is used only for declaration -> need to be fill
	for i := 0; i < 10; i++ {
		array1[i] = string(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("value ", i, " : ", array1[i], " memory address : ", &array1[i])
	}

}
