package workflow

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"../abstraction"
	"../decision"
	"../human"
	mylog "../log"
	"../send"
)

// Workflow contains tasks list of workflow definition.
type Workflow struct {
	ID            bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name          string        `json:"name" bson:"name"`
	State         string        `json:"state" bson:"state"`
	Tasks         []interface{} `json:"tasks" bson:"tasks"`
	ChannelsCount int           `json:"channelsCount" bson:"channelsCount"`
}

// NewWorkflow creates a new workflow definition.
func NewWorkflow(name string) *Workflow {
	wf := &Workflow{
		ID:    bson.NewObjectId(),
		Name:  name,
		State: "new",
		Tasks: make([]interface{}, 0),
	}
	return wf
}

func GetWorkflowById(id string) *Workflow {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//SELECT TABLE
	workflowsTable := session.DB("test").C("workflows")

	//CHECK IF STATE IS DONE
	result := Workflow{}
	err = workflowsTable.Find(bson.M{"_id": id}).One(&result)
	if err != nil && err.Error() != "not found" {
		log.Fatal(err)
	}

	return &result
}

func GetWorkflow(wf *Workflow) *Workflow {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//SELECT TABLE
	workflowsTable := session.DB("test").C("workflows")

	//CHECK IF STATE IS DONE
	result := Workflow{}
	err = workflowsTable.Find(bson.M{"_id": wf.ID}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

func SaveWorkflow(wf *Workflow) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//SELECT TABLE
	workflowsTable := session.DB("test").C("workflows")

	//INSERT
	err = workflowsTable.Insert(wf)
	if err != nil {
		panic(err)
	}
}

// AddTask add task with name.
func (wf *Workflow) AddTask(task interface{}) {
	wf.Tasks = append(wf.Tasks, task)
}

// Run defined workflow tasks.
func (wf *Workflow) Run() error {
	return wf.run(wf.Tasks)
}

func (wf *Workflow) run(tasks []interface{}) error {
	wf.State = "inprogress"

	for _, nt := range tasks {

		switch nttype := nt.(type) {
		case *mylog.LogTask:
			fmt.Println("LogTask - ", nttype)
			// RUN IT in separate goroutine
			go func(t abstraction.Task) {
				t.Execute()
			}(nttype)
		case *decision.BaseDecisionTask:
			fmt.Println("BaseDecisionTask - ", nttype)
			// RUN IT in separate goroutine
			go func(t abstraction.Task) {
				t.Execute()
			}(nttype)
		case *human.BaseHumanTask:
			fmt.Println("BaseHumanTask - ", nttype)
			// RUN IT in separate goroutine
			go func(t abstraction.Task) {
				t.Execute()
			}(nttype)
		case *send.SendEmailTask:
			fmt.Println("SendEmailTask -", nttype)
			// RUN IT in separate goroutine
			go func(t abstraction.Task) {
				t.Execute()
			}(nttype)
		case *send.SendSmsTask:
			fmt.Println("SendSmsTask -", nttype)
			// RUN IT in separate goroutine
			go func(t abstraction.Task) {
				t.Execute()
			}(nttype)
		case *send.SendMqttTask:
			fmt.Println("SendMqttTask -", nttype)
			// RUN IT in separate goroutine
			go func(t abstraction.Task) {
				t.Execute()
			}(nttype)
		case *send.SendRpcTask:
			fmt.Println("SendRpcTask -", nttype)
			// RUN IT in separate goroutine
			go func(t abstraction.Task) {
				t.Execute()
			}(nttype)
		default:
			fmt.Println("----default-----", nttype)
			// fmt.Println(nttype)
		}

	}

	wf.State = "completed"
	return nil
}
