package main

import (
	"fmt"
	"sync"
)

// ******************************************************
// RACE Conditions are bad things - avoid them
// go run -race main.go -> will help you
// you also must define how many goroutines must WaitGroup handle, before you use them
// ******************************************************

var c chan int
var wg sync.WaitGroup

func main() {

	c = make(chan int)

	wg.Add(2)

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
	wg.Done()
}

func waitAndCloseChannel() {
	wg.Wait()
	close(c)
}

func reader() {
	for n := range c {
		fmt.Println(n)
	}
}
