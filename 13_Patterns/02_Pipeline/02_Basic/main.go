package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
)

func main() {
	fmt.Println("START")

	topic := "s2315/temperature"

	// Set up channel on which to send signal notifications.
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	// Create an MQTT Client.
	cli := client.New(&client.Options{
		// Define the processing of the error handler.
		ErrorHandler: func(err error) {
			log.Println("4")
			fmt.Println(err)
		},
	})

	// Terminate the Client.
	defer cli.Terminate()

	// Connect to the MQTT Server.
	err2 := cli.Connect(&client.ConnectOptions{
		Network:  "tcp",
		Address:  "127.0.0.1:1883",
		ClientID: []byte("example-client5"),
	})
	if err2 != nil {
		log.Println("5")
		panic(err2)
	}

	// Subscribe to topics.
	err2 = cli.Subscribe(&client.SubscribeOptions{
		SubReqs: []*client.SubReq{
			&client.SubReq{
				TopicFilter: []byte(topic),
				QoS:         mqtt.QoS1,
				// Define the processing of the message handler.
				Handler: processMessage,
			},
		},
	})
	if err2 != nil {
		log.Println("6")
		panic(err2)
	}

	// Wait for receiving a signal.
	<-sigc

	// Disconnect the Network Connection.
	if err2 := cli.Disconnect(); err2 != nil {
		log.Println("7")
		panic(err2)
	}

	fmt.Println("END")
}

func processMessage(topicName, message []byte) {

	topic := string(topicName)
	senzorID := strings.Split(topic, "/")[0]
	measurement := strings.Split(topic, "/")[1]
	value := string(message)

	fmt.Println(topic)
	fmt.Println(senzorID)
	fmt.Println(measurement)
	fmt.Println(value)

	st1 := Stage1(value)
	st2 := Stage2(st1)
	st3 := Stage3(st2)

	fmt.Println(<-st3)
}

// Stage1 - some comment
func Stage1(value string) chan string {
	out := make(chan string)
	go func() {

		valueInt, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		if valueInt > 30 {
			out <- "HOT"
		}

		if valueInt < 10 {
			out <- "COLD"
		}

		close(out)
	}()
	return out
}

// Stage2 - some comment
func Stage2(in chan string) chan string {
	out := make(chan string)
	go func() {

		for n := range in {

			fmt.Println("Sending email")

			out <- n
		}

		close(out)
	}()

	go func() {

		for n := range in {

			fmt.Println("Sending sms")

			out <- n
		}

		close(out)
	}()

	go func() {

		for n := range in {

			fmt.Println("Publish to Twitter")

			out <- n
		}

		close(out)
	}()
	return out
}

// Stage3 - some comment
func Stage3(in chan string) chan string {
	out := make(chan string)
	go func() {

		for n := range in {

			if n == "HOT" {
				fmt.Println("Shine RED LED - via MQTT")
			} else if n == "COLD" {
				fmt.Println("Shine BLUE LED - via MQTT")
			}

			out <- n
		}

		close(out)
	}()
	return out
}
