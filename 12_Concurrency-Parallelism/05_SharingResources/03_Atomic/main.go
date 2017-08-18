package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
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
var counter int64

// ----------------------

func doSomething(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {

		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)

		//ATOMICALLY increment value about 1 ------
		atomic.AddInt64(&counter, 1)
		// ----------------------------------------

		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}
