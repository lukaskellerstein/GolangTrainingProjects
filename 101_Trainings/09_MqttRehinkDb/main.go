package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
	r "gopkg.in/gorethink/gorethink.v3"
)

//SenzorData - class for SenzorData object
type SenzorData struct {
	ID          string    `gorethink:"id,omitempty"`
	SenzorID    string    `gorethink:"senzorid"`
	Measurement string    `gorethink:"measurement"`
	Values      []string  `gorethink:"values"`
	Date        time.Time `gorethink:"date"`
}

var session *r.Session
var err error 

func main() {
	fmt.Println("START")

	//RethinkDB ----------------------------------------
	session, err = r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "test",
	})
	if err != nil {
		log.Fatalln(err)
	}

	//Recreate tables
	// err = r.DB("test").TableDrop("SenzorData").Exec(session)
	// err = r.DB("test").TableCreate("SenzorData").Exec(session)

	_, err = r.DB("test").Table("SenzorData").IndexCreate("senzorid").Run(session)
	_, err = r.DB("test").Table("SenzorData").IndexCreate("measurement").Run(session)
	if err != nil {
		log.Fatalln(err)
	}

	//MQTT ----------------------------------------

	// Set up channel on which to send signal notifications.
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	// Create an MQTT Client.
	cli := client.New(&client.Options{
		// Define the processing of the error handler.
		ErrorHandler: func(err error) {
			fmt.Println(err)
		},
	})

	// Terminate the Client.
	defer cli.Terminate()

	// Connect to the MQTT Server.
	err2 := cli.Connect(&client.ConnectOptions{
		Network:  "tcp",
		Address:  "192.168.1.234:1883",
		ClientID: []byte("example-client5"),
	})
	if err2 != nil {
		panic(err2)
	}

	// Subscribe to topics.
	err2 = cli.Subscribe(&client.SubscribeOptions{
		SubReqs: []*client.SubReq{
			&client.SubReq{
				TopicFilter: []byte("+/temperature"),
				QoS:         mqtt.QoS1,
				// Define the processing of the message handler.
				Handler: processMessage,
			},
			&client.SubReq{
				TopicFilter: []byte("+/humidity"),
				QoS:         mqtt.QoS1,
				// Define the processing of the message handler.
				Handler: processMessage,
			},
			&client.SubReq{
				TopicFilter: []byte("+/pir"),
				QoS:         mqtt.QoS1,
				// Define the processing of the message handler.
				Handler: processMessage,
			},
		},
	})
	if err2 != nil {
		panic(err2)
	}

	// Wait for receiving a signal.
	<-sigc

	// Disconnect the Network Connection.
	if err2 := cli.Disconnect(); err2 != nil {
		panic(err2)
	}

	// --------------------------------------------------

	fmt.Println("END")
}

func processMessage(topicName, message []byte) {

	topic := string(topicName)
	senzorID := strings.Split(topic, "/")[0]
	measurement := strings.Split(topic, "/")[1]
	value := string(message)

	nowTime := time.Now()

	newdata := new(SenzorData)
	newdata.SenzorID = senzorID
	newdata.Measurement = measurement
	newdata.Values = append(newdata.Values, value)
	newdata.Date = nowTime

	results2, err2 := r.DB("test").Table("SenzorData").Filter(func(row r.Term) r.Term {
		return row.Field("senzorid").Eq(senzorID)
	}).Filter(func(row r.Term) r.Term {
		return row.Field("measurement").Eq(measurement)
	}).Run(session)

	if err2 != nil {
		log.Fatal(err2)
	} else {

		if results2.IsNil() {
			// INSERT NEW ONE ----------------------------
			_, err := r.DB("test").Table("SenzorData").Insert(newdata).RunWrite(session)

			if err != nil {
				log.Fatal(err)
			}
			// -------------------------------------------
		} else {
			// UPDATE EXISTING ONE ----------------------------
			var response SenzorData
			err = results2.One(&response)

			if err != nil {
				log.Fatalln(err)
			}
			// fmt.Println(response)

			//CHANGE
			response.Values = append(response.Values, value)

			//UPDATE
			_, err3 := r.DB("test").Table("SenzorData").Update(response).Run(session)
			if err3 != nil {
				log.Fatalln(err3)
			}
			// ------------------------------------------------
		}

	}

	fmt.Println(string(topicName), string(message))
}
