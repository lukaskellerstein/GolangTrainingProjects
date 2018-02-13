package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

type TestObject struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
}

var mongodburl = "localhost"

//Shared connection to the DB
var session *mgo.Session
var err error

func main() {
	var err error
	if session, err = mgo.Dial(mongodburl); err != nil {
		log.Fatal(err)
	}
}

// Each call method = connecting to the DB
func saveToDB_One(obj TestObject) {
	session, err := mgo.Dial(mongodburl)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//SELECT TABLE
	senzordatatable := session.DB("test").C("table1")

	//INSERT
	senzordatatable.Insert(&obj)
}

// Each call method = connecting to the DB
func saveToDB_Two(obj *TestObject) {
	session, err := mgo.Dial(mongodburl)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//SELECT TABLE
	senzordatatable := session.DB("test").C("table2")

	//INSERT
	senzordatatable.Insert(obj)
}

// shared connecting to the DB
func saveToDB_Three(obj TestObject) {
	//SELECT TABLE
	senzordatatable := session.DB("test").C("table3")

	//INSERT
	senzordatatable.Insert(&obj)
}

// shared connecting to the DB
func saveToDB_Four(obj *TestObject) {
	//SELECT TABLE
	senzordatatable := session.DB("test").C("table3")

	//INSERT
	senzordatatable.Insert(obj)
}
