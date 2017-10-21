package main

import (
	"fmt"

	"log"

	"gopkg.in/mgo.v2/bson"
)

// Workflow contains tasks list of workflow definition.
type Workflow struct {
	ID     bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name   string        `json:"name" bson:"name"`
	State  string        `json:"state" bson:"state"`
	Tasks  []*Task       `json:"tasks" bson:"tasks"`
	logger *log.Logger
}

// NewWorkflow creates a new workflow definition.
func NewWorkflow(name string) *Workflow {
	return &Workflow{
		Name:  name,
		State: "new",
		Tasks: make([]*Task, 0),
	}
}

// AddTask add task with name.
func (wf *Workflow) AddTask(task Task) {
	wf.Tasks = append(wf.Tasks, &task)
}

// Run defined workflow tasks.
func (wf *Workflow) Run() error {
	return wf.run(wf.Tasks)
}

func (wf *Workflow) run(tasks []*Task) error {
	wf.State = "inprogress"
	for i, t := range tasks {
		wf.logger.Print(fmt.Sprintf("workflow: Start task: %v", tasks[i].GetName()))
		if err := t.Task.Execute(); err != nil {
			return err
		}
		wf.logger.Print(fmt.Sprintf("workflow: Complete task: %v", tasks[i].Name))
	}
	wf.State = "completed"
	return nil
}
