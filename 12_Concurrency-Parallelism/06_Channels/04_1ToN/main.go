package main

import (
	"fmt"
)

var c chan int
var done chan bool

func main() {

	n := 10
	c = make(chan int)
	done = make(chan bool)

	go writer()

	for i := 0; i < n; i++ {
		go reader()
	}

	for i := 0; i < n; i++ {
		<-done
	}
}

func writer() {
	for i := 0; i < 10000; i++ {
		c <- i
	}
	close(c)
}

func reader() {
	for n := range c {
		fmt.Println(n)
	}
	done <- true
}
