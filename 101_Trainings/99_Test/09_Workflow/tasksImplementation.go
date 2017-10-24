package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//**********************************
//TASK - Human Task
//**********************************
type someHumanTask struct {
	ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	State      string        `json:"state" bson:"state"`
	Value      string        `json:"value" bson:"value"`
	inChannel  chan string
	outChannel chan string
}

func NewHumanTask(name string) *someHumanTask {
	ht := &someHumanTask{
		ID:    bson.NewObjectId(),
		Name:  name,
		State: "new",
	}

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//SELECT TABLE
	humanTasksTable := session.DB("test").C("humanTasks")

	//INSERT
	err = humanTasksTable.Insert(ht)

	return ht
}

func (t *someHumanTask) SetID(id bson.ObjectId) {
	t.ID = id
}

func (t *someHumanTask) GetID() bson.ObjectId {
	return t.ID
}

func (t *someHumanTask) SetName(name string) {
	t.Name = name
}

func (t *someHumanTask) GetName() string {
	return t.Name
}

func (t *someHumanTask) SetState(state string) {
	t.State = state
}

func (t *someHumanTask) GetState() string {
	return t.State
}

func (t *someHumanTask) Execute() error {

	fmt.Println("someHumanTask1")

	for value := range t.inChannel {
		t.SetState("inprogress")

		session, err := mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
		defer session.Close()

		//SELECT TABLE
		humanTasksTable := session.DB("test").C("humanTasks")

		//CHECK IF STATE IS DONE
		result := someHumanTask{}
		err = humanTasksTable.Find(bson.M{"_id": t.ID}).One(&result)
		if err != nil {
			log.Fatal(err)
		}

		if result.State == "done" {
			fmt.Println("someHumanTask - " + t.Value)
			t.outChannel <- value
		}
	}

	fmt.Println("someHumanTask2")

	t.SetState("completed")
	return nil
}

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

	fmt.Println("extremeValueCheckTask1")

	for value := range t.inChannel {
		t.SetState("inprogress")
		t.Value = value
		fmt.Println("extremeValueCheckTask - " + t.Value)

		t.outChannel <- value
	}

	//SEM SE TO NEDOSTANE
	fmt.Println("extremeValueCheckTask2")

	//wg.Done()
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
	fmt.Println("sendEmailTask1")

	for value := range t.inChannel {
		t.SetState("inprogress")
		t.Value = value
		fmt.Println("sendEmailTask - " + t.Value)

		t.outChannel <- value
	}

	//SEM SE TO NEDOSTANE
	fmt.Println("sendEmailTask2")

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
	fmt.Println("sendSmsTask1")

	for value := range t.inChannel {
		t.SetState("inprogress")
		t.Value = value
		fmt.Println("sendSmsTask - " + t.Value)

		t.outChannel <- value
	}

	fmt.Println("sendSmsTask2")

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
	fmt.Println("twitterPostTask1")

	for value := range t.inChannel {
		t.SetState("inprogress")
		t.Value = value
		fmt.Println("twitterPostTask - " + t.Value)

		t.outChannel <- value
	}

	fmt.Println("twitterPostTask2")
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
	fmt.Println("sendToDatabase1")

	for value := range t.inChannel {
		t.SetState("inprogress")
		t.Value = value
		fmt.Println("sendToDatabase - " + t.Value)

		t.outChannel <- value
	}

	fmt.Println("sendToDatabase2")
	t.SetState("completed")
	return nil
}
