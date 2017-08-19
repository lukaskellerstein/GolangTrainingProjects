package main

import "fmt"
import "time"

var c chan int

func main() {

	c = make(chan int)

	go writer()
	go reader()

	// wrong way howto 'waiting' to end of goroutines
	// in future you will see better way howto do that
	time.Sleep(time.Second)
}

func writer() {
	for i := 0; i < 10; i++ {
		c <- i
	}
}

func reader() {
	for {
		fmt.Println(<-c)
	}
}
