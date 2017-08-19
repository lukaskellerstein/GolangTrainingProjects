package main

import "fmt"

func main() {
	channel1 := writerChannel()
	channel2 := chainedChannel(channel1)
	for value := range channel2 {
		fmt.Println(value)
	}
}

func writerChannel() chan int {
	outVar := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			outVar <- i
		}
		close(outVar)
	}()

	return outVar
}

func chainedChannel(inputChannel chan int) chan int {
	outVar := make(chan int)

	go func() {
		var sum int
		for value := range inputChannel {
			sum += value
		}
		outVar <- sum
		close(outVar)
	}()

	return outVar
}
