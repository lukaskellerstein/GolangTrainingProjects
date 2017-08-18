package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex

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

// ----------------------

func doSomething(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {

		// LOCK THESE RESOURCES
		mutex.Lock()

		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)

		// UNLOCK THESE RESOURCES
		mutex.Unlock()
	}
	wg.Done()
}
