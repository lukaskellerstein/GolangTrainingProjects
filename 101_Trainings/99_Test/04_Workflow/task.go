package main

import (
	"sync"

	multierror "github.com/hashicorp/go-multierror"

	"gopkg.in/mgo.v2/bson"
)

// Task represents task interface of workflow.
type Task interface {
	SetID(id bson.ObjectId)
	GetID() bson.ObjectId
	SetName(name string)
	GetName() string
	SetState(state string)
	GetState() string
	Execute(channel chan<- string) error
}

// type baseTask struct {
// 	ID    bson.ObjectId `json:"_id" bson:"_id,omitempty"`
// 	Name  string        `json:"name" bson:"name"`
// 	State string        `json:"state" bson:"state"`
// 	Task  Task          `json:"task" bson:"task"`
// }

// ParallelTask represents parallel task on workflow.
type ParallelTask struct {
	ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	State      string        `json:"state" bson:"state"`
	Tasks      []Task        `json:"tasks" bson:"tasks"`
	inChannel  chan string
	outChannel chan string
}

// NewParallelTask creates a parallel task by task list.
func NewParallelTask(name string, inchannel chan string, outchannel chan string) *ParallelTask {
	return &ParallelTask{
		Name:       name,
		State:      "new",
		Tasks:      make([]Task, 0),
		inChannel:  inchannel,
		outChannel: outchannel,
	}
}

// AddTask add parallel task with name
func (pt *ParallelTask) AddTask(task Task) {
	pt.Tasks = append(pt.Tasks, task)
}

func (t *ParallelTask) SetID(id bson.ObjectId) {
	t.ID = id
}

func (t *ParallelTask) GetID() bson.ObjectId {
	return t.ID
}

func (t *ParallelTask) SetName(name string) {
	t.Name = name
}

func (t *ParallelTask) GetName() string {
	return t.Name
}

func (t *ParallelTask) SetState(state string) {
	t.State = state
}

func (t *ParallelTask) GetState() string {
	return t.State
}

// Execute implement Task.Execute.
func (pt *ParallelTask) Execute(channel chan<- string) error {
	pt.State = "inprogress"
	errChan := make(chan error)
	var wg sync.WaitGroup

	for _, nt := range pt.Tasks {
		wg.Add(1)
		go func(t Task) {
			if err := t.Execute(channel); err != nil {
				errChan <- err
			}
			wg.Done()
		}(nt)
	}

	resultChan := make(chan error)
	go func() {
		var result *multierror.Error
		for err := range errChan {
			result = multierror.Append(result, err)
		}
		resultChan <- result.ErrorOrNil()
	}()

	wg.Wait()
	close(errChan)

	pt.State = "completed"
	return <-resultChan
}
