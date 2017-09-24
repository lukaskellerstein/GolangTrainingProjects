package main

import (
	"fmt"
)

func main() {

	for index := 0; index < 10; index++ {
		fmt.Println(index)
	}

	a := 10

	if a == 10 {
		fmt.Println("a = 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else {
		fmt.Println("a < 10")
	}

	switch a {
	default:
		fmt.Println("always run")
	case 1:
		fmt.Println("a = 1")
	case 2:
		fmt.Println("a = 2")
	case 3:
		fmt.Println("a = 3")
	}
}
