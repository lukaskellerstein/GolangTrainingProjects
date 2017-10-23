package main

import (
	"fmt"

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
	return &Workflow{
		Name:  name,
		State: "new",
		Tasks: make([]Task, 0),
	}
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

	// myChan := make(chan string)
	// errChan := make(chan error)
	//var wg sync.WaitGroup

	for _, nt := range tasks {
		//wg.Add(1)

		fmt.Println("workflow: Start task: " + nt.GetName())

		go func(t Task) {
			t.Execute()
			//wg.Done()
		}(nt)

		fmt.Println("workflow: Run task: " + nt.GetName())
	}

	// resultChan := make(chan error)
	// go func() {
	// 	var result *multierror.Error
	// 	for err := range errChan {
	// 		result = multierror.Append(result, err)
	// 	}
	// 	resultChan <- result.ErrorOrNil()
	// }()

	// wg.Wait()
	// close(errChan)

	//??? - poradne pochopit, proc to spoustim v nove goroutine
	// go func() {
	// 	wg.Wait()
	// 	close(myChan)
	// }()

	// for message := range myChan {
	// 	fmt.Println(message)
	// }

	// go func() {
	// 	wg.Wait()
	// }()

	wf.State = "completed"
	return nil
}
