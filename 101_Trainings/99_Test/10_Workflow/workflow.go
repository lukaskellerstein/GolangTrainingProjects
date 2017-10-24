package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Workflow contains tasks list of workflow definition.
type Workflow struct {
	ID    bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name  string        `json:"name" bson:"name"`
	State string        `json:"state" bson:"state"`
	Tasks []Task        `json:"tasks" bson:"tasks"`
}

// NewWorkflow creates a new workflow definition.
func NewWorkflow(name string) *Workflow {
	wf := &Workflow{
		ID:    bson.NewObjectId(),
		Name:  name,
		State: "new",
		Tasks: make([]Task, 0),
	}

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

	return wf
}

// AddTask add task with name.
func (wf *Workflow) AddTask(task Task) {
	wf.Tasks = append(wf.Tasks, task)
}

// Run defined workflow tasks.
func (wf *Workflow) Run() error {
	return wf.run(wf.Tasks)
}

func (wf *Workflow) run(tasks []Task) error {
	wf.State = "inprogress"

	for _, nt := range tasks {
		fmt.Println("workflow: Start task: " + nt.GetName())

		go func(t Task) {
			t.Execute()
		}(nt)

		fmt.Println("workflow: Run task: " + nt.GetName())
	}

	wf.State = "completed"
	return nil
}
