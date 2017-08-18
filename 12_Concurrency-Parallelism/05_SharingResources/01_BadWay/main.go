package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	// concurrency, but no parallelism
	go doSomething("One: ")
	go doSomething("Two: ")
	wg.Wait()
	// The program is waiting until both goroutines are done
}

// SHARED RESOURCE ------
var counter int

// resource will be edited at same time multiple resource,
// but when set value first function,
// this value will be override second function
// ----------------------

func doSomething(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}
