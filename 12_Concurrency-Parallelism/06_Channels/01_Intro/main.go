package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	// wrong way howto 'waiting' to end of goroutines
	// in future you will see better way howto do that
	time.Sleep(time.Second)
}
