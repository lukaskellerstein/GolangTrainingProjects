package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//docker run -d -p 27017:27017 -t cellar.hub.mongodb

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

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//SELECT TABLE
	senzordatatable := session.DB("test").C("senzordatatable")

	//INSERT
	err = senzordatatable.Insert(&SenzorData{bson.NewObjectId(), "Senzor1", "Temperature", []string{"34", "22", "10"}, time.Now()},
		&SenzorData{bson.NewObjectId(), "Senzor2", "Temperature", []string{"54", "12", "11"}, time.Now()})

	if err != nil {
		panic(err)
	}

	fmt.Println("END")
}
