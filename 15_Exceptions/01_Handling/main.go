package main

import (
	"errors"
	"fmt"
)

func main() {
	defer recoverPanic()

	err := errors.New("TEST PANIC")

	panic(err)

}

//Error handling
func recoverPanic() {
	if rec := recover(); rec != nil {
		err := rec.(error)

		//low-level exception logging
		fmt.Println(err.Error())
		// fmt.Println("[PANIC] - " + err.Error())

		// os.Exit(1)
	}
}
