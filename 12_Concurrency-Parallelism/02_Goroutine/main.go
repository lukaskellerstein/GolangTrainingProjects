package main

import "fmt"

func main() {
	// concurrency, but no parallelism
	go doSomething1()
	go doSomething2()
	// The program is closed immediately after start goroutines, so this
	// mean, the program doesn't matter if goroutines ends alright or not
}

func doSomething1() {
	for i := 0; i < 10000; i++ {
		// We don't see anything because goroutines runs in separate "OS thread"
		// and the program (main function) is closed immediately after start this goroutine
		fmt.Printf("doSomething1 - %v \n", i)
	}
}

func doSomething2() {
	for i := 0; i < 10000; i++ {
		// We don't see anything because goroutines runs in separate "OS thread"
		// and the program (main function) is closed immediately after start this goroutine
		fmt.Printf("doSomething2 - %v \n", i)
	}
}
