package main

import (
	"encoding/json"
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

			//normal task
			ch1out := make(chan string)
			wf.AddTask(&extremeValueCheckTask{BaseTask: BaseTask{Type: "extremeValueCheckTask", Name: "myTask1", State: "new", ID: bson.NewObjectId(), inChannel: workflowIn, outChannel: ch1out}, MinValue: "10", MaxValue: "20"})

			//parallel task
			ch2out := make(chan string)
			pt := NewParallelTask("Parallel1", ch1out, ch2out)
			pt.AddTask(&sendEmailTask{BaseTask: BaseTask{Type: "sendEmailTask", Name: "myTask2", State: "new", ID: bson.NewObjectId()}, EmailAddress: "someuser@gmail.com"})
			pt.AddTask(&sendSmsTask{BaseTask{Type: "sendSmsTask", Name: "myTask3", State: "new", ID: bson.NewObjectId()}})
			pt.AddTask(&twitterPostTask{BaseTask{Type: "twitterPostTask", Name: "myTask4", State: "new", ID: bson.NewObjectId()}})
			wf.AddTask(pt)

			//human task
			ch3out := make(chan string)
			ht := NewHumanTask("ht1")
			ht.UserID = "user21"
			ht.ResolvedTime = "27-10-2017 15:43:00"
			ht.inChannel = ch2out
			ht.outChannel = ch3out
			wf.AddTask(ht)

			//decision task
			ch4out := make(chan string)
			ch5out := make(chan string)
			ch6out := make(chan string)
			chas := make([]chan string, 3)
			chas[0] = ch4out
			chas[1] = ch5out
			chas[2] = ch6out

			dt := NewDecisionTask("Decision1", ch3out, chas)
			wf.AddTask(dt)

			ch7out := make(chan string)
			wf.AddTask(&sendEmailTask{BaseTask: BaseTask{Type: "sendEmailTask", Name: "myTask5", State: "new", ID: bson.NewObjectId(), inChannel: ch4out, outChannel: ch7out}, EmailAddress: "someuser@gmail.com"})
			wf.AddTask(&sendSmsTask{BaseTask{Type: "sendSmsTask", Name: "myTask6", State: "new", ID: bson.NewObjectId(), inChannel: ch5out, outChannel: ch7out}})
			wf.AddTask(&twitterPostTask{BaseTask{Type: "twitterPostTask", Name: "myTask7", State: "new", ID: bson.NewObjectId(), inChannel: ch6out, outChannel: ch7out}})

			//normal task
			wf.AddTask(&sendToDatabase{BaseTask: BaseTask{Type: "sendToDatabase", Name: "myTask8", inChannel: ch7out, outChannel: workflowOut}, DatabaseName: "testDatabase"})

			// ---------------------------
			SaveWorkflow(wf)
			// ---------------------------

			//TEST ____________________________________________
			//_________________________________________________
			wfram := Workflow{}
			isUsedin := false
			lastUsedChannels := make([]chan string, 3)
			mychannels := make([]chan string, 100)
			wfdb := GetWorkflow(wf)
			for _, nt := range wfdb.Tasks {

				baseTaskVariable := nt.(bson.M)["basetask"]
				concreteTaskVariable := nt.(bson.M)

				newinChannel := make(chan string)
				if isUsedin == false {
					isUsedin = true
					newinChannel = workflowIn
				} else if isUsedin == true {
					if len(lastUsedChannels) == 0 {
						// ???
					} else {
						//get last element
						newinChannel = lastUsedChannels[len(lastUsedChannels)-1]
						//delete last element
						lastUsedChannels = lastUsedChannels[:len(lastUsedChannels)-1]

						lastUsedChannels.re
					}
				}
				newoutChannel := make(chan string)

				if concreteTaskVariable["type"] == "BaseParallelTask" {
					asdf := BaseParallelTask{}
					bodyBytes, _ := json.Marshal(concreteTaskVariable)
					json.Unmarshal(bodyBytes, &asdf)
					// fmt.Println(asdf)

					// asdf.inChannel = newinChannel
					// asdf.outChannel = newoutChannel

					wfram.AddTask(asdf)
				}

				if concreteTaskVariable["type"] == "BaseDecisionTask" {
					asdf := BaseDecisionTask{}
					bodyBytes, _ := json.Marshal(concreteTaskVariable)
					json.Unmarshal(bodyBytes, &asdf)
					// fmt.Println(asdf)

					// asdf.inChannel = newinChannel
					// asdf.outChannel = newoutChannel

					wfram.AddTask(asdf)
				}

				//BaseTask
				asdf := BaseTask{}
				bodyBytes, _ := json.Marshal(baseTaskVariable)
				json.Unmarshal(bodyBytes, &asdf)
				// fmt.Println(asdf)

				//COCNRETE TASK
				if asdf.Type == "extremeValueCheckTask" {
					asdf2 := extremeValueCheckTask{BaseTask: asdf}
					bodyBytes, _ := json.Marshal(concreteTaskVariable)
					json.Unmarshal(bodyBytes, &asdf2)

					// fmt.Println(asdf2)
					// fmt.Println(asdf2.ID)
					// fmt.Println(asdf2.Name)
					// fmt.Println(asdf2.MinValue)
					// fmt.Println(asdf2.MaxValue)

					// asdf2.inChannel = newinChannel
					// asdf2.outChannel = newoutChannel

					wfram.AddTask(asdf2)
				} else if asdf.Type == "someHumanTask" {
					asdf2 := someHumanTask{BaseTask: asdf}
					bodyBytes, _ := json.Marshal(concreteTaskVariable)
					json.Unmarshal(bodyBytes, &asdf2)

					// fmt.Println(asdf2)
					// fmt.Println(asdf2.ID)
					// fmt.Println(asdf2.Name)
					// fmt.Println(asdf2.UserID)
					// fmt.Println(asdf2.ResolvedTime)

					// asdf2.inChannel = newinChannel
					// asdf2.outChannel = newoutChannel

					wfram.AddTask(asdf2)
				} else if asdf.Type == "sendEmailTask" {
					asdf2 := sendEmailTask{BaseTask: asdf}
					bodyBytes, _ := json.Marshal(concreteTaskVariable)
					json.Unmarshal(bodyBytes, &asdf2)

					// fmt.Println(asdf2)
					// fmt.Println(asdf2.ID)
					// fmt.Println(asdf2.Name)
					// fmt.Println(asdf2.EmailAddress)

					// asdf2.inChannel = newinChannel
					// asdf2.outChannel = newoutChannel

					wfram.AddTask(asdf2)
				} else if asdf.Type == "sendSmsTask" {
					asdf2 := sendSmsTask{BaseTask: asdf}
					bodyBytes, _ := json.Marshal(concreteTaskVariable)
					json.Unmarshal(bodyBytes, &asdf2)

					// fmt.Println(asdf2)
					// fmt.Println(asdf2.ID)
					// fmt.Println(asdf2.Name)

					// asdf2.inChannel = newinChannel
					// asdf2.outChannel = newoutChannel

					wfram.AddTask(asdf2)
				} else if asdf.Type == "twitterPostTask" {
					asdf2 := twitterPostTask{BaseTask: asdf}
					bodyBytes, _ := json.Marshal(concreteTaskVariable)
					json.Unmarshal(bodyBytes, &asdf2)

					// fmt.Println(asdf2)
					// fmt.Println(asdf2.ID)
					// fmt.Println(asdf2.Name)

					// asdf2.inChannel = newinChannel
					// asdf2.outChannel = newoutChannel

					wfram.AddTask(asdf2)
				} else if asdf.Type == "sendToDatabase" {
					asdf2 := sendToDatabase{BaseTask: asdf}
					bodyBytes, _ := json.Marshal(concreteTaskVariable)
					json.Unmarshal(bodyBytes, &asdf2)

					// fmt.Println(asdf2)
					// fmt.Println(asdf2.ID)
					// fmt.Println(asdf2.Name)
					// fmt.Println(asdf2.DatabaseName)

					// asdf2.inChannel = newinChannel
					// asdf2.outChannel = newoutChannel

					wfram.AddTask(asdf2)
				}

				// v := reflect.ValueOf(nt)
				// t := v.Type()
				// k := t.Kind()
				// fmt.Println(v)
				// fmt.Println(t)
				// fmt.Println(k)
			}
			//_________________________________________________

			tempvar := "ddd"
			fmt.Println(tempvar)

			//wfram.Run()
			wfdb.Run()

			// wf.Run()
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
