package main

import (
	"fmt"
)

type SomeObject struct {
	text string
}

func main() {
	fmt.Println("START")

	//*******************************
	// SERIOVE ZPRACOVANI
	//*******************************
	newChan1 := startPipeline()
	newChan2 := writeIntoChannel(newChan1)
	newChan3 := writeIntoChannel(newChan2)
	newChan4 := writeIntoChannel(newChan3)
	newChan5 := writeIntoChannel(newChan4)

	fmt.Println(<-newChan5)

	fmt.Println("END")
}

func startPipeline() chan SomeObject {
	out := make(chan SomeObject)

	go func() {
		out <- SomeObject{text: "START_pipeline|"}
		close(out)
	}()

	return out
}

func writeIntoChannel(channel chan SomeObject) chan SomeObject {
	out := make(chan SomeObject)

	go func() {
		for obj := range channel {
			obj.text += "AAAA|"
			out <- obj
		}
		close(out)
	}()

	return out
}
