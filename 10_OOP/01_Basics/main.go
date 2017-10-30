package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"
)

var workflowIn chan string
var workflowOut chan string

func main() {
	fmt.Println("START")

	// Set up channel on which to send signal notifications.
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	workflowIn = make(chan string)
	workflowOut = make(chan string)

	//-------------------------------------------------------------------
	//-------------------------------------------------------------------

	workflowName := "Workflow1213"
	RunWorkflow1(workflowName)

	//-------------------------------------------------------------------
	//-------------------------------------------------------------------
	// each 2 seconds send message

	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(2 * time.Second)
			randomNumber := random(1, 100)
			workflowIn <- strconv.Itoa(randomNumber)
		}
		close(workflowIn)
	}()

	//-------------------------------------------------------------------
	//-------------------------------------------------------------------

	// go func() {
	// 	for value := range workflowOut {
	// 		fmt.Println("OUT > " + value)
	// 	}
	// 	close(workflowOut)
	// }()

	for value := range workflowOut {
		fmt.Println("OUT > " + value)
	}
	//close(workflowOut)

	// Wait for receiving a signal.
	<-sigc

	fmt.Println("END")
}

//HELPER
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
