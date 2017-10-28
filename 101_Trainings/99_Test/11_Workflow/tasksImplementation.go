package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//**********************************
//TASK - Human Task
//**********************************
type someHumanTask struct {
	BaseTask
	UserID       string `json:"userID" bson:"userID"`
	ResolvedTime string `json:"resolvedTime" bson:"resolvedTime"`
}

func NewHumanTask(name string) *someHumanTask {
	ht := &someHumanTask{
		BaseTask: BaseTask{
			ID:         bson.NewObjectId(),
			Name:       name,
			State:      "new",
			Value:      "",
			inChannel:  nil,
			outChannel: nil,
		},
		UserID:       "asdf",
		ResolvedTime: "asdf",
	}

	return ht
}

func GetHumanTaskState(ht *someHumanTask) string {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//SELECT TABLE
	workflowsTable := session.DB("test").C("workflows")

	result00 := Workflow{}
	err = workflowsTable.Find(nil).Select(bson.M{"name": 1}).One(&result00)
	if err != nil {
		log.Fatal(err)
	}

	result01 := Workflow{}
	err = workflowsTable.Find(nil).One(&result01)
	if err != nil {
		log.Fatal(err)
	}

	//CHECK IF STATE IS DONE
	result := someHumanTask{}
	err = workflowsTable.Find(bson.M{"tasks._id": ht.ID}).Select(bson.M{"tasks": 1}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	result2 := someHumanTask{}
	err = workflowsTable.Find(bson.M{"tasks._id": ht.ID}).Select(bson.M{"tasks.$": 1}).One(&result2)
	if err != nil {
		log.Fatal(err)
	}

	// result3 := Workflow{}
	// err = humanTasksTable.Find(bson.M{"tasks._id": ht.ID}).Select(bson.M{"tasks": 1}).One(&result3)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	result4 := Workflow{}
	err = workflowsTable.Find(bson.M{"tasks._id": ht.ID}).Select(bson.M{"tasks.$": 1}).One(&result4)
	if err != nil {
		log.Fatal(err)
	}

	return result.State
}

func (t *someHumanTask) Execute() error {

	fmt.Println("someHumanTask1")

	for value := range t.inChannel {
		t.State = "inprogress"

		result := GetHumanTaskState(t)

		if result == "done" {
			fmt.Println("someHumanTask - " + t.Value)
			t.outChannel <- value
		}
	}

	//SEM SE TO NEDOSTANE ???
	fmt.Println("someHumanTask2")

	t.State = "completed"
	return nil
}

//**********************************
//TASK - Is there some extreme value ?
//**********************************
type extremeValueCheckTask struct {
	BaseTask
	MinValue string `json:"minValue" bson:"minValue"`
	MaxValue string `json:"maxValue" bson:"maxValue"`
}

// func UpdateTask(t *extremeValueCheckTask) {
// 	session, err := mgo.Dial("localhost")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer session.Close()

// 	//SELECT TABLE
// 	workflowsTable := session.DB("test").C("workflows")

// 	// Update
// 	colQuerier := bson.M{"tasks.name": "Senzor1"}
// 	change := bson.M{"$set": bson.M{"date": time.Now()}}

// 	err = workflowsTable.Update(colQuerier, change)
// 	if err != nil {
// 		panic(err)
// 	}

// 	//SELECT TABLE
// 	humanTasksTable := session.DB("test").C("humanTasks")

// 	//CHECK IF STATE IS DONE
// 	result := extremeValueCheckTask{}
// 	err = humanTasksTable.Find(bson.M{"_id": t.ID}).One(&result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func (t *extremeValueCheckTask) Execute() error {

	fmt.Println("extremeValueCheckTask1")

	for value := range t.inChannel {

		if t.State == "new" {
			t.State = "inprogress"
		}

		t.State = "inprogress"
		t.Value = value

		// //SELECT TABLE
		// humanTasksTable := session.DB("test").C("humanTasks")

		// //CHECK IF STATE IS DONE
		// result := someHumanTask{}
		// err = humanTasksTable.Find(bson.M{"_id": t.ID}).One(&result)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// if result.State == "done" {
		// 	fmt.Println("someHumanTask - " + t.Value)
		// 	t.outChannel <- value
		// }

		fmt.Println("extremeValueCheckTask - " + t.Value)

		t.outChannel <- value
	}

	//SEM SE TO NEDOSTANE ??? vim proc
	fmt.Println("extremeValueCheckTask2")

	//wg.Done()
	t.State = "completed"
	return nil
}

//**********************************
//TASK - Send email
//**********************************
type sendEmailTask struct {
	BaseTask
	EmailAddress string `json:"emailAddress" bson:"emailAddress"`
}

func (t *sendEmailTask) Execute() error {
	fmt.Println("sendEmailTask1")

	for value := range t.inChannel {
		t.State = "inprogress"
		t.Value = value
		fmt.Println("sendEmailTask - " + t.Value)

		t.outChannel <- value
	}

	//SEM SE TO NEDOSTANE ???
	fmt.Println("sendEmailTask2")

	t.State = "completed"
	return nil
}

//**********************************
//TASK - Send sms
//**********************************
type sendSmsTask struct {
	BaseTask
}

func (t *sendSmsTask) Execute() error {
	fmt.Println("sendSmsTask1")

	for value := range t.inChannel {
		t.State = "inprogress"
		t.Value = value
		fmt.Println("sendSmsTask - " + t.Value)

		t.outChannel <- value
	}

	//SEM SE TO NEDOSTANE ???
	fmt.Println("sendSmsTask2")

	t.State = "completed"
	return nil
}

//**********************************
//TASK - Send Twitter post
//**********************************
type twitterPostTask struct {
	BaseTask
}

func (t *twitterPostTask) Execute() error {
	fmt.Println("twitterPostTask1")

	for value := range t.inChannel {
		t.State = "inprogress"
		t.Value = value
		fmt.Println("twitterPostTask - " + t.Value)

		t.outChannel <- value
	}

	//SEM SE TO NEDOSTANE ???
	fmt.Println("twitterPostTask2")

	t.State = "completed"
	return nil
}

//**********************************
//TASK - Save to the Database
//**********************************
type sendToDatabase struct {
	BaseTask
	DatabaseName string `json:"databaseName" bson:"databaseName"`
}

func (t *sendToDatabase) Execute() error {
	fmt.Println("sendToDatabase1")

	for value := range t.inChannel {
		t.State = "inprogress"
		t.Value = value
		fmt.Println("sendToDatabase - " + t.Value)

		t.outChannel <- value
	}

	//SEM SE TO NEDOSTANE ???
	fmt.Println("sendToDatabase2")

	t.State = "completed"
	return nil
}
