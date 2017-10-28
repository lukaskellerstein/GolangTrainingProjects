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

			ch1out := make(chan string)
			wf.AddTask(&extremeValueCheckTask{BaseTask: BaseTask{Name: "extremeValueCheckTask", State: "new", ID: bson.NewObjectId(), inChannel: workflowIn, outChannel: ch1out}, MinValue: "10", MaxValue: "20"})

			ch2out := make(chan string)
			wf.AddTask(&sendEmailTask{BaseTask: BaseTask{Name: "sendEmailTask", State: "new", ID: bson.NewObjectId(), inChannel: ch1out, outChannel: ch2out}, EmailAddress: "someuser@gmail.com"})

			ch3out := make(chan string)
			wf.AddTask(&sendSmsTask{BaseTask{Name: "sendSmsTask", State: "new", ID: bson.NewObjectId(), inChannel: ch2out, outChannel: ch3out}})

			ch4out := make(chan string)
			wf.AddTask(&twitterPostTask{BaseTask{Name: "twitterPostTask", State: "new", ID: bson.NewObjectId(), inChannel: ch3out, outChannel: ch4out}})

			ch5out := make(chan string)
			ht := NewHumanTask("ht1")
			ht.UserID = "user21"
			ht.ResolvedTime = "27-10-2017 15:43:00"
			ht.inChannel = ch4out
			ht.outChannel = ch5out
			wf.AddTask(ht)

			wf.AddTask(&sendToDatabase{BaseTask: BaseTask{Name: "sendToDatabase", inChannel: ch5out, outChannel: workflowOut}, DatabaseName: "testDatabase"})

			// ---------------------------
			SaveWorkflow(wf)
			// ---------------------------

			//TEST ____________________________________________
			//_________________________________________________
			asdf := GetWorkflow(wf)
			for _, nt := range asdf.Tasks {

				// switch nttype := nt.(type) {
				// case *someHumanTask:
				// 	fmt.Println("someHumanTask*")
				// case *extremeValueCheckTask:
				// 	fmt.Println("extremeValueCheckTask*")
				// case *sendEmailTask:
				// 	fmt.Println("sendEmailTask*")
				// case *sendSmsTask:
				// 	fmt.Println("sendSmsTask*")
				// case *twitterPostTask:
				// 	fmt.Println("twitterPostTask*")
				// case *sendToDatabase:
				// 	fmt.Println("sendToDatabase*")
				// case someHumanTask:
				// 	fmt.Println("someHumanTask")
				// case extremeValueCheckTask:
				// 	fmt.Println("extremeValueCheckTask")
				// case sendEmailTask:
				// 	fmt.Println("sendEmailTask")
				// case sendSmsTask:
				// 	fmt.Println("sendSmsTask")
				// case twitterPostTask:
				// 	fmt.Println("twitterPostTask")
				// case sendToDatabase:
				// 	fmt.Println("sendToDatabase")
				// case bson.M:
				// 	fmt.Println("bson.M")
				// default:
				// 	fmt.Println("----default-----")
				// 	fmt.Println(nttype)
				// }

				// aaa0 := nt.(bson.M)
				// fmt.Println(aaa0)

				// aaa2 := nt.(bson.M)["minValue"]
				// fmt.Println(aaa2)

				// aaa1 := nt.(bson.M)["basetask"]
				// fmt.Println(aaa1)
				// // aaa1a := aaa1.(BaseTask)
				// // BaseTask(aaa1)

				baseTaskVariable := nt.(bson.M)["basetask"]
				concreteTaskVariable := nt.(bson.M)

				//BaseTask
				asdf := BaseTask{}
				bodyBytes, _ := json.Marshal(baseTaskVariable)
				json.Unmarshal(bodyBytes, &asdf)
				fmt.Println(asdf)

				//COCNRETE TASK
				if asdf.Name == "extremeValueCheckTask" {
					asdf2 := extremeValueCheckTask{BaseTask: asdf}
					bodyBytes, _ := json.Marshal(concreteTaskVariable)
					json.Unmarshal(bodyBytes, &asdf2)

					fmt.Println(asdf2)
					fmt.Println(asdf2.ID)
					fmt.Println(asdf2.Name)
					fmt.Println(asdf2.MinValue)
					fmt.Println(asdf2.MaxValue)
				} else if asdf.Name == "someHumanTask" {
					asdf2 := someHumanTask{BaseTask: asdf}
					bodyBytes, _ := json.Marshal(concreteTaskVariable)
					json.Unmarshal(bodyBytes, &asdf2)

					fmt.Println(asdf2)
					fmt.Println(asdf2.ID)
					fmt.Println(asdf2.Name)
					fmt.Println(asdf2.UserID)
					fmt.Println(asdf2.ResolvedTime)
				} else if asdf.Name == "sendEmailTask" {
					asdf2 := sendEmailTask{BaseTask: asdf}
					bodyBytes, _ := json.Marshal(concreteTaskVariable)
					json.Unmarshal(bodyBytes, &asdf2)

					fmt.Println(asdf2)
					fmt.Println(asdf2.ID)
					fmt.Println(asdf2.Name)
					fmt.Println(asdf2.EmailAddress)
				} else if asdf.Name == "sendSmsTask" {
					asdf2 := sendSmsTask{BaseTask: asdf}
					bodyBytes, _ := json.Marshal(concreteTaskVariable)
					json.Unmarshal(bodyBytes, &asdf2)

					fmt.Println(asdf2)
					fmt.Println(asdf2.ID)
					fmt.Println(asdf2.Name)
				} else if asdf.Name == "twitterPostTask" {
					asdf2 := twitterPostTask{BaseTask: asdf}
					bodyBytes, _ := json.Marshal(concreteTaskVariable)
					json.Unmarshal(bodyBytes, &asdf2)

					fmt.Println(asdf2)
					fmt.Println(asdf2.ID)
					fmt.Println(asdf2.Name)
				} else if asdf.Name == "sendToDatabase" {
					asdf2 := sendToDatabase{BaseTask: asdf}
					bodyBytes, _ := json.Marshal(concreteTaskVariable)
					json.Unmarshal(bodyBytes, &asdf2)

					fmt.Println(asdf2)
					fmt.Println(asdf2.ID)
					fmt.Println(asdf2.Name)
					fmt.Println(asdf2.DatabaseName)
				}

				// v := reflect.ValueOf(nt)
				// t := v.Type()
				// k := t.Kind()
				// fmt.Println(v)
				// fmt.Println(t)
				// fmt.Println(k)
			}
			//_________________________________________________

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
