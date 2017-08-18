package main

import "fmt"

func main() {
	// no concurrency neither parallelism
	doSomething1()
	doSomething2()
}

func doSomething1() {
	for i := 0; i < 100; i++ {
		fmt.Printf("doSomething1 - %v \n", i)
	}
}

func doSomething2() {
	for i := 0; i < 100; i++ {
		fmt.Printf("doSomething2 - %v \n", i)
	}
}
