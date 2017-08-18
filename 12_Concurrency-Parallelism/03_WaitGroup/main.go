package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	// concurrency, but no parallelism
	go doSomething1()
	go doSomething2()
	wg.Wait()
	// The program is waiting until both goroutines are done
}

func doSomething1() {
	for i := 0; i < 100; i++ {
		fmt.Printf("doSomething1 - %v \n", i)

		// this is there only for example of long running task
		time.Sleep(30 * time.Millisecond)
	}
	wg.Done()
}

func doSomething2() {
	for i := 0; i < 100; i++ {
		fmt.Printf("doSomething2 - %v \n", i)

		// this is there only for example of long running task
		time.Sleep(200 * time.Millisecond)
	}
	wg.Done()
}
