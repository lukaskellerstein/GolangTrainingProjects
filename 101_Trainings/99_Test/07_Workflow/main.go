package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
)

var workflowIn chan string
var workflowOut chan string

func main() {
	fmt.Println("START")

	workflowIn = make(chan string)
	workflowOut = make(chan string)

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

	//*******************************
	// WORKFLOW
	//*******************************
	go func() {
		wf := NewWorkflow("Workflow1")

		ch1out := make(chan string)
		wf.AddTask(&extremeValueCheckTask{Name: "extremeValueCheckTask", inChannel: workflowIn, outChannel: ch1out})

		ch2out := make(chan string)
		ch3out := make(chan string)
		ch4out := make(chan string)
		var chas [3]chan string
		chas[0] = ch2out
		chas[1] = ch3out
		chas[2] = ch4out

		pt := NewDecisionTask("Decision1", ch1out, chas)
		wf.AddTask(pt)

		ch5out := make(chan string)

		wf.AddTask(&sendEmailTask{Name: "sendEmailTask", inChannel: ch2out, outChannel: ch5out})
		wf.AddTask(&sendSmsTask{Name: "sendSmsTask", inChannel: ch3out, outChannel: ch5out})
		wf.AddTask(&twitterPostTask{Name: "twitterPostTask", inChannel: ch4out, outChannel: ch5out})

		wf.AddTask(&sendToDatabase{Name: "sendToDatabase", inChannel: ch5out, outChannel: workflowOut})

		wf.Run()
	}()

	//naliti dat do prvniho channelu
	workflowIn <- value

	for value := range workflowOut {
		fmt.Println("OUT > " + value)
	}

}
