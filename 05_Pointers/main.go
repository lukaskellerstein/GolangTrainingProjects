package main

import (
	"fmt"
)

func main() {

	// referencing
	temp1 := 42
	temp2 := temp1  //only copy a value
	temp3 := &temp1 //copy memory address

	fmt.Println(temp1)
	fmt.Println(temp2)
	fmt.Println(temp3)

	temp1++

	fmt.Println(temp1)
	fmt.Println(temp2)
	fmt.Println(temp3)

	temp2++

	fmt.Println(temp1)
	fmt.Println(temp2)
	fmt.Println(temp3)

	temp2++
	temp3 = &temp2

	fmt.Println(temp1)
	fmt.Println(temp2)
	fmt.Println(temp3) //this is pointer to value 44

	// Dereferencing = get back a value
	fmt.Println(*temp3) //a proof that, this will be 44

	// Using pointers
	*temp3 = 999 //set the memory address to value 999

	fmt.Println(temp1)
	fmt.Println(temp2)
	fmt.Println(temp3)

}
