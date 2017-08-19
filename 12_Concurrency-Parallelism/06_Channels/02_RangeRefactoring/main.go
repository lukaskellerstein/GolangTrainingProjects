package main

import (
	"fmt"
	"time"
)

var c chan int

func main() {
	c = make(chan int)

	go writer()
	go reader()

	// wrong way howto 'waiting' to end of goroutines
	time.Sleep(time.Second)
}

func writer() {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func reader() {
	// this code block is waiting and range through incoming values,
	// until somebody close channel 'c'
	for n := range c {
		fmt.Println(n)
	}
}
