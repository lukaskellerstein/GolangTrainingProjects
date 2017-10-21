package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

//**********************************
//TASK - Is there some extreme value ?
//**********************************
type extremeValueCheckTask struct {
	iD    bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	name  string        `json:"name" bson:"name"`
	state string        `json:"state" bson:"state"`
	Value string        `json:"value" bson:"value"`
}

func (t *extremeValueCheckTask) SetID(id bson.ObjectId) {
	t.iD = id
}

func (t *extremeValueCheckTask) GetID() bson.ObjectId {
	return t.iD
}

func (t *extremeValueCheckTask) SetName(name string) {
	t.name = name
}

func (t *extremeValueCheckTask) GetName() string {
	return t.name
}

func (t *extremeValueCheckTask) SetState(state string) {
	t.state = state
}

func (t *extremeValueCheckTask) GetState() string {
	return t.state
}

func (t *extremeValueCheckTask) Execute() error {
	t.SetState("inprogress")

	t.Value = string(rand.Intn(50))

	i, err := strconv.Atoi(t.Value)
	if err != nil {
		return err
	}

	if i < 10 {
		fmt.Println("extremeValueCheckTask = x < 10")
	}

	if i > 30 {
		fmt.Println("extremeValueCheckTask = x > 30")
	}
	t.SetState("completed")
	return nil
}

//**********************************
//TASK - Send email
//**********************************
type sendEmailTask struct {
	// Name  string `json:"name" bson:"name"`
	// State string `json:"state" bson:"state"`

	Value string `json:"value" bson:"value"`
}

func (t *sendEmailTask) Execute() error {
	// t.State = "inprogress"
	fmt.Println("sendEmailTask - " + t.Value)
	// t.State = "completed"
	return nil
}

//**********************************
//TASK - Send sms
//**********************************
type sendSmsTask struct {
	// Name  string `json:"name" bson:"name"`
	// State string `json:"state" bson:"state"`

	Value string `json:"value" bson:"value"`
}

func (t *sendSmsTask) Execute() error {
	// t.State = "inprogress"
	// fmt.Println("sendSmsTask - " + t.Name)
	// t.State = "completed"
	return nil
}

//**********************************
//TASK - Send Twitter post
//**********************************
type twitterPostTask struct {
	// Name  string `json:"name" bson:"name"`
	// State string `json:"state" bson:"state"`

	Value string `json:"value" bson:"value"`
}

func (t *twitterPostTask) Execute() error {
	// t.State = "inprogress"
	// fmt.Println("twitterPostTask - " + t.Name)
	// t.State = "completed"
	return nil
}

//**********************************
//TASK - Save to the Database
//**********************************
type sendToDatabase struct {
	// Name  string `json:"name" bson:"name"`
	// State string `json:"state" bson:"state"`

	Value string `json:"value" bson:"value"`
}

func (t *sendToDatabase) Execute() error {
	// t.State = "inprogress"
	// fmt.Println("sendToDatabase - " + t.Name)
	// t.State = "completed"
	return nil
}
