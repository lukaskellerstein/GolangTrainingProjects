package main

import (
	"fmt"
	"sync"
)

type SomeObject struct {
	text string
}

var wg sync.WaitGroup

func main() {
	fmt.Println("START")

	//*******************************
	// PARARELNI ZPRACOVANI
	//*******************************
	myChan := make(chan SomeObject)

	go writeIntoChannel(myChan)
	wg.Add(1)
	go writeIntoChannel(myChan)
	wg.Add(1)
	go writeIntoChannel(myChan)
	wg.Add(1)

	//??? - poradne pochopit, proc to spoustim v nove goroutine
	go func() {
		wg.Wait()
		close(myChan)
	}()

	for message := range myChan {
		fmt.Println(message)
	}

	fmt.Println("END")
}

func writeIntoChannel(channel chan<- SomeObject) {

	obj := SomeObject{text: "asdf"}
	channel <- obj
	wg.Done()
}
