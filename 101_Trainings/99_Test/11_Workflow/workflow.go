package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Workflow contains tasks list of workflow definition.
type Workflow struct {
	ID    bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name  string        `json:"name" bson:"name"`
	State string        `json:"state" bson:"state"`
	Tasks []BaseTask    `json:"tasks" bson:"tasks"`
}

// NewWorkflow creates a new workflow definition.
func NewWorkflow(name string) *Workflow {
	wf := &Workflow{
		ID:    bson.NewObjectId(),
		Name:  name,
		State: "new",
		Tasks: make([]BaseTask, 0),
	}

	//SaveWorkflow(wf)

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
func (wf *Workflow) AddTask(task BaseTask) {
	wf.Tasks = append(wf.Tasks, task)
}

// Run defined workflow tasks.
func (wf *Workflow) Run() error {
	return wf.run(wf.Tasks)
}

func (wf *Workflow) run(tasks []BaseTask) error {
	wf.State = "inprogress"

	// myChan := make(chan string)
	// errChan := make(chan error)
	//var wg sync.WaitGroup

	for _, nt := range tasks {
		//wg.Add(1)

		fmt.Println("workflow: Start task: " + nt.Name)

		go func(t BaseTask) {
			t.Execute()
			//wg.Done()
		}(nt)

		fmt.Println("workflow: Run task: " + nt.Name)
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
