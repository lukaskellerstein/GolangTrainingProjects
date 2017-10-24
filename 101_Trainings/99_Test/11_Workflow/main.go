package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
	"gopkg.in/mgo.v2/bson"
)

var workflowIn chan string
var workflowOut chan string

var triggerTopic string
var workflowId string

func main() {
	fmt.Println("START")

	workflowIn = make(chan string)
	workflowOut = make(chan string)

	// Set up channel on which to send signal notifications.
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	//*******************************
	// TRIGGER
	//*******************************
	topic := "s2315/temperature"

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

	//*******************************
	// WORKFLOW
	//*******************************
	workflowId = "59ef036911bbacaebc995fc0"

	isExistWorkflow := GetWorkflowById(workflowId)

	if isExistWorkflow.ID != "" {
		go func() {
			isExistWorkflow.Run()
		}()
	} else {
		go func() {
			wf := NewWorkflow("Workflow22")

			// PIPELINE ------------------

			ch1out := make(chan string)
			wf.AddTask(&extremeValueCheckTask{BaseTask{Name: "extremeValueCheckTask", State: "new", ID: bson.NewObjectId(), inChannel: workflowIn, outChannel: ch1out}})

			ch2out := make(chan string)
			wf.AddTask(&sendEmailTask{BaseTask{Name: "sendEmailTask", State: "new", ID: bson.NewObjectId(), inChannel: ch1out, outChannel: ch2out}})

			ch3out := make(chan string)
			wf.AddTask(&sendSmsTask{BaseTask{Name: "sendSmsTask", State: "new", ID: bson.NewObjectId(), inChannel: ch2out, outChannel: ch3out}})

			ch4out := make(chan string)
			wf.AddTask(&twitterPostTask{BaseTask{Name: "twitterPostTask", State: "new", ID: bson.NewObjectId(), inChannel: ch3out, outChannel: ch4out}})

			ch5out := make(chan string)
			ht := NewHumanTask("ht1")
			ht.inChannel = ch4out
			ht.outChannel = ch5out
			wf.AddTask(ht)

			wf.AddTask(&sendToDatabase{BaseTask{Name: "sendToDatabase", inChannel: ch5out, outChannel: workflowOut}})

			// ---------------------------
			SaveWorkflow(wf)
			// ---------------------------

			wf.Run()
		}()
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
	fmt.Println(senzorID + "/" + measurement + " - " + value)

	//naliti dat do prvniho channelu
	workflowIn <- value

	for value := range workflowOut {
		fmt.Println("OUT > " + value)
	}

}
