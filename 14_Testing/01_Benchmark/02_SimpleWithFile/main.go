package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		return
	}

	result := ReadAndReplace(file)

	fmt.Println(result)
}
