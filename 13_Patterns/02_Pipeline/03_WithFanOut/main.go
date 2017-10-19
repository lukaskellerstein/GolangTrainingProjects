package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"

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

// Stage1 *************************************
// Stage1 - some comment
// Stage1 *************************************
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

// Stage2 *************************************
// Stage2 - some comment
// Stage2 *************************************
func Stage2(in chan string) chan string {
	out := make(chan string)
	go func() {

		done := make(chan int)
		defer close(done)

		// Distribute the sq work across two goroutines that both read from in.
		c1 := SendMail(done, in)
		c2 := SendSms(done, in)
		c3 := SendTwitter(done, in)

		for n := range merge(done, c1, c2, c3) {
			fmt.Println(n, string("asdf"))
		}

		close(out)
	}()
	return out
}

// SendMail step
func SendMail(done chan int, in chan string) chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for n := range in {
			fmt.Println(n, string("mail1"))

			select {
			default:
				fmt.Println("SendMail - default")
			case out <- n:
				fmt.Println(n, string("mail2"))
			case <-done:
				fmt.Print("sq -- done")
				return
			}
		}
	}()
	return out
}

// SendSms step
func SendSms(done chan int, in chan string) chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for n := range in {
			fmt.Println(n, string("sms1"))

			select {
			default:
				fmt.Println("SendSms - default")
			case out <- n:
				fmt.Println(n, string("sms2"))
			case <-done:
				fmt.Print("sq -- done")
				return
			}
		}
	}()
	return out
}

// SendTwitter step
func SendTwitter(done chan int, in chan string) chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for n := range in {
			fmt.Println(n, string("twitter1"))

			select {
			default:
				fmt.Println("SendTwitter - default")
			case out <- n:
				fmt.Println(n, string("twitter2"))
			case <-done:
				fmt.Print("sq -- done")
				return
			}
		}
	}()
	return out
}

func merge(done chan int, cs ...chan string) chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c chan string) {
		defer wg.Done()
		for n := range c {
			fmt.Println(n, string("merge1"))

			select {
			default:
				fmt.Println("merge - default")
			case out <- n:
				fmt.Println(n, string("merge2"))
			case <-done:
				fmt.Print("done")
				return
			}
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Stage3 *************************************
// Stage3 - some comment
// Stage3 *************************************
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
