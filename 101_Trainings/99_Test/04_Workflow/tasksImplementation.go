package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//**********************************
//TASK - Is there some extreme value ?
//**********************************
type extremeValueCheckTask struct {
	ID    bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name  string        `json:"name" bson:"name"`
	State string        `json:"state" bson:"state"`
	Value string        `json:"value" bson:"value"`
}

func (t *extremeValueCheckTask) SetID(id bson.ObjectId) {
	t.ID = id
}

func (t *extremeValueCheckTask) GetID() bson.ObjectId {
	return t.ID
}

func (t *extremeValueCheckTask) SetName(name string) {
	t.Name = name
}

func (t *extremeValueCheckTask) GetName() string {
	return t.Name
}

func (t *extremeValueCheckTask) SetState(state string) {
	t.State = state
}

func (t *extremeValueCheckTask) GetState() string {
	return t.State
}

func (t *extremeValueCheckTask) Execute(channel chan<- string) error {
	t.SetState("inprogress")

	i := rand.Intn(50)

	if i < 10 {
		fmt.Println("extremeValueCheckTask = x < 10")
	}

	if i > 30 {
		fmt.Println("extremeValueCheckTask = x > 30")
	}

	time.Sleep(2 * time.Second)

	t.Value = strconv.Itoa(i)

	t.SetState("completed")
	return nil
}

//**********************************
//TASK - Send email
//**********************************
type sendEmailTask struct {
	ID    bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name  string        `json:"name" bson:"name"`
	State string        `json:"state" bson:"state"`
	Value string        `json:"value" bson:"value"`
}

func (t *sendEmailTask) SetID(id bson.ObjectId) {
	t.ID = id
}

func (t *sendEmailTask) GetID() bson.ObjectId {
	return t.ID
}

func (t *sendEmailTask) SetName(name string) {
	t.Name = name
}

func (t *sendEmailTask) GetName() string {
	return t.Name
}

func (t *sendEmailTask) SetState(state string) {
	t.State = state
}

func (t *sendEmailTask) GetState() string {
	return t.State
}

func (t *sendEmailTask) Execute(channel chan<- string) error {
	t.SetState("inprogress")
	time.Sleep(2 * time.Second)
	i := rand.Intn(50)
	t.Value = strconv.Itoa(i)
	fmt.Println("sendEmailTask - " + t.Value)
	t.SetState("completed")
	return nil
}

//**********************************
//TASK - Send sms
//**********************************
type sendSmsTask struct {
	ID    bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name  string        `json:"name" bson:"name"`
	State string        `json:"state" bson:"state"`
	Value string        `json:"value" bson:"value"`
}

func (t *sendSmsTask) SetID(id bson.ObjectId) {
	t.ID = id
}

func (t *sendSmsTask) GetID() bson.ObjectId {
	return t.ID
}

func (t *sendSmsTask) SetName(name string) {
	t.Name = name
}

func (t *sendSmsTask) GetName() string {
	return t.Name
}

func (t *sendSmsTask) SetState(state string) {
	t.State = state
}

func (t *sendSmsTask) GetState() string {
	return t.State
}

func (t *sendSmsTask) Execute(channel chan<- string) error {
	t.SetState("inprogress")
	time.Sleep(2 * time.Second)
	i := rand.Intn(50)
	t.Value = strconv.Itoa(i)
	fmt.Println("sendSmsTask - " + t.Value)
	t.SetState("completed")
	return nil
}

//**********************************
//TASK - Send Twitter post
//**********************************
type twitterPostTask struct {
	ID    bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name  string        `json:"name" bson:"name"`
	State string        `json:"state" bson:"state"`
	Value string        `json:"value" bson:"value"`
}

func (t *twitterPostTask) SetID(id bson.ObjectId) {
	t.ID = id
}

func (t *twitterPostTask) GetID() bson.ObjectId {
	return t.ID
}

func (t *twitterPostTask) SetName(name string) {
	t.Name = name
}

func (t *twitterPostTask) GetName() string {
	return t.Name
}

func (t *twitterPostTask) SetState(state string) {
	t.State = state
}

func (t *twitterPostTask) GetState() string {
	return t.State
}

func (t *twitterPostTask) Execute(channel chan<- string) error {
	t.SetState("inprogress")
	time.Sleep(2 * time.Second)
	i := rand.Intn(50)
	t.Value = strconv.Itoa(i)
	fmt.Println("twitterPostTask - " + t.Value)
	t.SetState("completed")
	return nil
}

//**********************************
//TASK - Save to the Database
//**********************************
type sendToDatabase struct {
	ID    bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name  string        `json:"name" bson:"name"`
	State string        `json:"state" bson:"state"`
	Value string        `json:"value" bson:"value"`
}

func (t *sendToDatabase) SetID(id bson.ObjectId) {
	t.ID = id
}

func (t *sendToDatabase) GetID() bson.ObjectId {
	return t.ID
}

func (t *sendToDatabase) SetName(name string) {
	t.Name = name
}

func (t *sendToDatabase) GetName() string {
	return t.Name
}

func (t *sendToDatabase) SetState(state string) {
	t.State = state
}

func (t *sendToDatabase) GetState() string {
	return t.State
}

func (t *sendToDatabase) Execute(channel chan<- string) error {
	t.SetState("inprogress")
	time.Sleep(2 * time.Second)
	i := rand.Intn(50)
	t.Value = strconv.Itoa(i)
	fmt.Println("sendToDatabase - " + t.Value)
	t.SetState("completed")
	return nil
}
