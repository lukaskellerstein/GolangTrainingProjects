package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AppDatabase interface {
	GetText() string
}

func NewMongoDatabase() AppDatabase {
	return MongoDatabase{}
}

type MongoDatabase struct{}

type TextData struct {
	ID   bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Text string        `json:"text" bson:"text"`
}

func (MongoDatabase) GetText() string {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//SELECT TABLE
	textsdatatable := session.DB("test").C("texts")

	//SELECT
	result := TextData{}
	err = textsdatatable.Find(nil).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result.Text
}
