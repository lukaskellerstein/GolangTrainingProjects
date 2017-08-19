package main

import (
	"fmt"
)

// ******************************************************
// RACE Conditions are bad things - avoid them
// go run -race main.go -> will help you
// you can also use Semaphore pattern, which is pure solution by using only channels
// Main thing of this pattern is using another channel named 'done' which receive state of other goroutines
// ******************************************************

var c chan int
var done chan bool

func main() {

	c = make(chan int)
	done = make(chan bool)

	// ***************************************************
	// It doesn't matter on order of runs goroutine
	// it will always be good
	// ***************************************************

	go writer()
	go writer()

	go waitAndCloseChannel()
	go reader()
}

func writer() {
	for i := 0; i < 10; i++ {
		c <- i
	}
	done <- true
}

func waitAndCloseChannel() {
	<-done
	<-done
	close(c)
}

func reader() {
	for n := range c {
		fmt.Println(n)
	}
}
