package main

import "fmt"

var globalVariable = 5

func main() {

	var localVariable = 6

	fmt.Println(globalVariable)
	fmt.Println(localVariable)
}
