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

func main() {

	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
