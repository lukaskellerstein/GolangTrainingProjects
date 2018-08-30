package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"
)

//IN-OUT CHANNELS ------------------
var someChannel chan string
var closeChannel chan string

func main() {

	// Set up channel on which to send signal notifications.
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	someChannel = make(chan string)
	closeChannel = make(chan string)

	//autogenerate number send to the channel
	go func() {
	loop:
		for {
			time.Sleep(time.Duration(1) * time.Second)
			randomNumberFloat := rand.Float64() * 1000

			select {
			case <-closeChannel:
				break loop // has to be named, because "break" applies to the select otherwise
			default:
				//do nothing
			}

			//send value to the channel
			someChannel <- strconv.FormatFloat(randomNumberFloat, 'E', -1, 64)
		}
	}()

	go func() {
		for value := range someChannel {
			fmt.Println(value)
		}
	}()

	go func() {
		time.Sleep(time.Duration(10) * time.Second)
		closeChannel <- "close all"
		close(someChannel)
	}()

	// Wait for receiving a signal.
	<-sigc

}
