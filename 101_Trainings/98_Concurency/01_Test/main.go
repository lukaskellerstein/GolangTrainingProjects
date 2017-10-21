package main

import (
	"fmt"
)

type SomeObject struct {
	text string
}

func main() {
	fmt.Println("START")

	myChan := make(chan SomeObject)

	go writeIntoChannel(myChan)
	go writeIntoChannel(myChan)
	go writeIntoChannel(myChan)

	//TEST if it will be bad, when we have more writes then reads
	//RESULT this can't be readed
	//go writeIntoChannel(myChan)

	fmt.Println((<-myChan).text)
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)

	//TEST if it will be bad, when we have more reads then writes
	//RESULT deadlock
	//fmt.Println(<-myChan)

	fmt.Println("END")
}

func writeIntoChannel(channel chan<- SomeObject) {

	obj := SomeObject{text: "asdf"}
	channel <- obj

}
