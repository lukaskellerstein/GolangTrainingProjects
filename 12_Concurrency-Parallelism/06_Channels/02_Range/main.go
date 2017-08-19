package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	// this code block is waiting and range through incoming values,
	// until somebody close channel 'c'
	for n := range c {
		fmt.Println(n)
	}
}
