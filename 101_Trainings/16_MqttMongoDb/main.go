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

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//SenzorData - class for SenzorData object
type SenzorData struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	SenzorID    string        `json:"senzorid" bson:"senzorid"`
	Measurement string        `json:"measurement" bson:"measurement"`
	Values      []string      `json:"values" bson:"values"`
	Date        time.Time     `json:"date" bson:"date"`
}

func main() {
	fmt.Println("START")

	//--------------------------------------------------
	//MongoDB ----------------------------------------
	//--------------------------------------------------

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//check if table exist --------------------------
	// IS NOT NEEDED, mongo create collection on demand
	// isExistSenzorDataTable := false
	// names, err := session.DB("test").CollectionNames()
	// if err != nil {
	// 	// Handle error
	// 	log.Printf("Failed to get coll names: %v", err)
	// 	return
	// }

	// // Simply search in the names slice, e.g.
	// for _, name := range names {
	// 	if name == "collectionToCheck" {
	// 		isExistSenzorDataTable = true
	// 		break
	// 	}
	// }

	// //Recreate tables ------------------------
	// if isExistSenzorDataTable == false {
	// 	//SELECT TABLE
	// 	senzordatatable := session.DB("test").C("senzordatatable")
	// }
	//--------------------------------------------------

	//SELECT TABLE
	senzordatatable := session.DB("test").C("senzordatatable")

	//Create indexes ------------------------
	index := mgo.Index{
		Name:       "senzoridIndex",
		Key:        []string{"senzorid"},
		Unique:     false,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	//CHECK IF EXIST > DELETE
	indexes, err := senzordatatable.Indexes()
	if err != nil {
		panic(err)
	}
	for _, indexEx := range indexes {
		err = senzordatatable.DropIndex(indexEx.Key...)
		if err != nil {
			fmt.Println(err)
		}
	}

	//senzordatatable.DropIndexName("senzoridIndex")

	//CREATE
	if err := senzordatatable.EnsureIndex(index); err != nil {
		panic(err)
	}

	//--------------------------------------------------
	//MQTT ---------------------------------------------
	//--------------------------------------------------

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

	//***************************
	//***************************

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	result := SenzorData{}
	err = session.DB("test").C("senzordatatable").Find(
		bson.M{
			"senzorid":    senzorID,
			"measurement": measurement,
		},
	).One(&result)

	if err != nil && err.Error() != "not found" {
		log.Println("8")
		log.Fatal(err)
	}

	if result.SenzorID == "" {

		// INSERT NEW ONE ----------------------------
		err = session.DB("test").C("senzordatatable").Insert(&newdata)

		if err != nil {
			panic(err)
		}
		// -------------------------------------------
	} else {
		// UPDATE EXISTING ONE ----------------------------
		colQuerier := bson.M{
			"senzorid":    senzorID,
			"measurement": measurement,
		}
		change := bson.M{"$push": bson.M{"values": value}}
		err = session.DB("test").C("senzordatatable").Update(colQuerier, change)
		if err != nil {
			panic(err)
		}

		// ------------------------------------------------
	}

	fmt.Println(string(topicName), string(message))
}
