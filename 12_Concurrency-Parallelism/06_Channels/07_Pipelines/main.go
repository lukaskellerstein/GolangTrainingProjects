package main

import (
	"fmt"
)

func main() {
	fmt.Println("START")

	fmt.Println("END")
}

// DODELAT podle https://rclayton.silvrback.com/pipelines-in-golang

func generateData(topic string) {

}

func subscribeData(topic string) chan string {
	chan1 := make(chan string)

	return chan1
}

func saveData(input <-chan string) {
	for n := range input {
		fmt.Println(n)
	}
}
