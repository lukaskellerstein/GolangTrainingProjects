package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

//**********************************
//TASK - Is there some extreme value ?
//**********************************
type extremeValueCheckTask struct {
	ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	State      string        `json:"state" bson:"state"`
	Value      string        `json:"value" bson:"value"`
	inChannel  chan string
	outChannel chan string
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

func (t *extremeValueCheckTask) Execute() error {

	fmt.Println("extremeValueCheckTask")

	for value := range t.inChannel {
		t.SetState("inprogress")
		t.Value = value
		fmt.Println("extremeValueCheckTask - " + t.Value)

		t.outChannel <- value
	}

	//SEM SE TO NEDOSTANE
	fmt.Println("extremeValueCheckTask")

	//wg.Done()
	t.SetState("completed")
	return nil
}

func (t *extremeValueCheckTask) ExecuteParallel(value string) error {
	t.SetState("inprogress")
	t.Value = value
	fmt.Println("extremeValueCheckTask - " + t.Value)
	t.SetState("completed")
	return nil
}

//**********************************
//TASK - Send email
//**********************************
type sendEmailTask struct {
	ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	State      string        `json:"state" bson:"state"`
	Value      string        `json:"value" bson:"value"`
	inChannel  chan string
	outChannel chan string
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

func (t *sendEmailTask) Execute() error {
	for value := range t.inChannel {
		t.SetState("inprogress")
		t.Value = value
		fmt.Println("sendEmailTask - " + t.Value)

		t.outChannel <- value
	}

	t.SetState("completed")
	return nil
}

func (t *sendEmailTask) ExecuteParallel(value string) error {
	t.SetState("inprogress")
	t.Value = value
	fmt.Println("sendEmailTask - " + t.Value)
	t.SetState("completed")
	return nil
}

//**********************************
//TASK - Send sms
//**********************************
type sendSmsTask struct {
	ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	State      string        `json:"state" bson:"state"`
	Value      string        `json:"value" bson:"value"`
	inChannel  chan string
	outChannel chan string
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

func (t *sendSmsTask) Execute() error {
	for value := range t.inChannel {
		t.SetState("inprogress")
		t.Value = value
		fmt.Println("sendSmsTask - " + t.Value)

		t.outChannel <- value
	}

	t.SetState("completed")
	return nil
}

func (t *sendSmsTask) ExecuteParallel(value string) error {
	t.SetState("inprogress")
	t.Value = value
	fmt.Println("sendSmsTask - " + t.Value)
	t.SetState("completed")
	return nil
}

//**********************************
//TASK - Send Twitter post
//**********************************
type twitterPostTask struct {
	ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	State      string        `json:"state" bson:"state"`
	Value      string        `json:"value" bson:"value"`
	inChannel  chan string
	outChannel chan string
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

func (t *twitterPostTask) Execute() error {
	for value := range t.inChannel {
		t.SetState("inprogress")
		t.Value = value
		fmt.Println("twitterPostTask - " + t.Value)

		t.outChannel <- value
	}

	t.SetState("completed")
	return nil
}

func (t *twitterPostTask) ExecuteParallel(value string) error {
	t.SetState("inprogress")
	t.Value = value
	fmt.Println("twitterPostTask - " + t.Value)
	t.SetState("completed")
	return nil
}

//**********************************
//TASK - Save to the Database
//**********************************
type sendToDatabase struct {
	ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	State      string        `json:"state" bson:"state"`
	Value      string        `json:"value" bson:"value"`
	inChannel  chan string
	outChannel chan string
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

func (t *sendToDatabase) Execute() error {
	for value := range t.inChannel {
		t.SetState("inprogress")
		t.Value = value
		fmt.Println("sendToDatabase - " + t.Value)

		t.outChannel <- value
	}

	t.SetState("completed")
	return nil
}

func (t *sendToDatabase) ExecuteParallel(value string) error {
	t.SetState("inprogress")
	t.Value = value
	fmt.Println("sendToDatabase - " + t.Value)
	t.SetState("completed")
	return nil
}
