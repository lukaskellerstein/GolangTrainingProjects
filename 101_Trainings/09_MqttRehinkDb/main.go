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

//docker run -d -p 28015:28015 -p 29015:29015 -p 8080:8080 -t cellar.hub.rethinkdb

//SenzorData - class for SenzorData object
type SenzorData struct {
	ID          string    `gorethink:"id,omitempty"`
	SenzorID    string    `gorethink:"senzorid"`
	Measurement string    `gorethink:"measurement"`
	Values      []string  `gorethink:"values"`
	Date        time.Time `gorethink:"date"`
}

func main() {
	fmt.Println("START")

	//--------------------------------------------------
	//RethinkDB ----------------------------------------
	//--------------------------------------------------

	//***************************
	fmt.Println("open session1")
	session1 := connectToDB()
	//***************************

	//Recreate tables if not exists -----------
	var response []interface{}
	res, err := r.DB("test").TableList().Run(session1)

	//***************************
	fmt.Println("close session1")
	closeConnectionToDB(session1)
	//***************************

	//QQQQ - proc tam davam "&" - Memory address a ne cely objekt ?
	err = res.All(&response)
	if err != nil {
		log.Println("2")
		fmt.Println(err)
	}

	//check if table exist
	isExistSenzorDataTable := false
	for _, db := range response {
		if db == "SenzorData" {
			isExistSenzorDataTable = true
		}
	}

	//Recreate tables ------------------------
	if isExistSenzorDataTable == false {
		// err = r.DB("test").TableDrop("SenzorData").Exec(session)

		//***************************
		fmt.Println("open session2")
		session2 := connectToDB()
		//***************************

		err = r.DB("test").TableCreate("SenzorData").Exec(session2)

		//***************************
		fmt.Println("close session2")
		closeConnectionToDB(session2)
		//***************************
	}

	//Create indexes ------------------------
	// _, err = r.DB("test").Table("SenzorData").IndexCreate("senzorid").Run(session)
	// _, err = r.DB("test").Table("SenzorData").IndexCreate("measurement").Run(session)
	if err != nil {
		log.Println("3")
		log.Fatalln(err)
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
	fmt.Println("open session3")
	session3 := connectToDB()
	//***************************

	results2, err2 := r.DB("test").Table("SenzorData").Filter(func(row r.Term) r.Term {
		return row.Field("senzorid").Eq(senzorID)
	}).Filter(func(row r.Term) r.Term {
		return row.Field("measurement").Eq(measurement)
	}).Run(session3)

	var response SenzorData
	if results2.IsNil() == false {

		err := results2.One(&response)

		if err != nil {
			log.Println("10")
			log.Fatalln(err)
		}
	}

	//***************************
	fmt.Println("close session3")
	closeConnectionToDB(session3)
	//***************************

	if err2 != nil {
		log.Println("8")
		log.Fatal(err2)
	} else {

		if response.ID == "" {

			//***************************
			fmt.Println("open session4")
			session4 := connectToDB()
			//***************************

			// INSERT NEW ONE ----------------------------
			_, err := r.DB("test").Table("SenzorData").Insert(newdata).RunWrite(session4)

			//***************************
			fmt.Println("close session4")
			closeConnectionToDB(session4)
			//***************************

			if err != nil {
				log.Println("9")
				log.Fatal(err)
			}
			// -------------------------------------------
		} else {
			// UPDATE EXISTING ONE ----------------------------

			//CHANGE
			response.Values = append(response.Values, value)

			//***************************
			fmt.Println("open session5")
			session5 := connectToDB()
			//***************************

			//UPDATE
			_, err3 := r.DB("test").Table("SenzorData").Update(response).Run(session5)

			//***************************
			fmt.Println("close session5")
			closeConnectionToDB(session5)
			//***************************

			if err3 != nil {
				log.Println("11")
				log.Fatalln(err3)
			}

			// ------------------------------------------------
		}

	}

	fmt.Println(string(topicName), string(message))
}

func connectToDB() *r.Session {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "test",
	})
	if err != nil {
		log.Println("1")
		log.Fatalln(err)
	}
	fmt.Println("open OK")
	return session
}
func closeConnectionToDB(session *r.Session) {
	session.Close()
	fmt.Println("close OK")
}
