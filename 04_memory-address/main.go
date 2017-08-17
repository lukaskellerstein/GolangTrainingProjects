package main

import "fmt"

func main() {
	someTest := "There is some test text"

	fmt.Println("in variable someTest is: ", someTest)
	fmt.Println("and it is save in memory address: ", &someTest)

	// I dont know howto get back string
	fmt.Printf("if i try translate it back to text: %d \n", &someTest)
	fmt.Printf("if i try translate it back to text: %s \n", &someTest)
	fmt.Printf("if i try translate it back to text: %p \n", &someTest)
}
