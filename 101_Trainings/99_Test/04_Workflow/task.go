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
	Execute() error
}

// type baseTask struct {
// 	ID    bson.ObjectId `json:"_id" bson:"_id,omitempty"`
// 	Name  string        `json:"name" bson:"name"`
// 	State string        `json:"state" bson:"state"`
// 	Task  Task          `json:"task" bson:"task"`
// }

// ParallelTask represents parallel task on workflow.
type ParallelTask struct {
	tasks []*Task
}

// NewParallelTask creates a parallel task by task list.
func NewParallelTask() *ParallelTask {
	return &ParallelTask{tasks: make([]*Task, 0)}
}

// AddTask add parallel task with name
func (pt *ParallelTask) AddTask(task Task) {
	pt.tasks = append(pt.tasks, &task)
}

// Execute implement Task.Execute.
func (pt *ParallelTask) Execute() error {
	errChan := make(chan error)
	var wg sync.WaitGroup

	for _, nt := range pt.tasks {
		wg.Add(1)
		go func(t Task) {
			if err := t.Execute(); err != nil {
				errChan <- err
			}
			wg.Done()
		}(*nt)
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

	return <-resultChan
}
