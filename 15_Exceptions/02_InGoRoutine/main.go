package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	defer recoverPanic()

	// Set up channel on which to send signal notifications.
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	go func() {
		defer recoverPanic()

		var exceptionCount = 0

		for {
			time.Sleep(1 * time.Second)

			if exceptionCount == 30 {
				errTest := errors.New("TEST PANIC time repeater")
				panic(errTest)
			}
			exceptionCount++

		}
	}()

	// Wait for receiving a signal.
	<-sigc
}

//Error handling
func recoverPanic() {
	if rec := recover(); rec != nil {
		err := rec.(error)

		//low-level exception logging
		fmt.Println(err.Error())
		// fmt.Println("[PANIC] - " + err.Error())

		// os.Exit(1)
	}
}
