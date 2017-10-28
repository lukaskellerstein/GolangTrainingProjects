package main

import (
	"fmt"
	"strconv"
	"sync"

	multierror "github.com/hashicorp/go-multierror"
	"gopkg.in/mgo.v2/bson"
)

// Task represents task interface of workflow.
type Task interface {
	Execute() error
	ExecuteParallel(value string) error
}

type BaseTask struct {
	ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	Type       string        `json:"type" bson:"type"`
	State      string        `json:"state" bson:"state"`
	Value      string        `json:"value" bson:"value"`
	inChannel  chan string
	outChannel chan string
}

// default empty implementation
func (t *BaseTask) Execute() error {
	// do nothing
	return nil
}

//*********************************************
// Base Parallel Task
// BaseParallelTask represents parallel task on workflow.
//*********************************************
type BaseParallelTask struct {
	ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	Type       string        `json:"type" bson:"type"`
	State      string        `json:"state" bson:"state"`
	Tasks      []interface{} `json:"tasks" bson:"tasks"`
	inChannel  chan string
	outChannel chan string
}

// NewParallelTask creates a parallel task by task list.
func NewParallelTask(name string, inchannel chan string, outchannel chan string) *BaseParallelTask {
	return &BaseParallelTask{
		Name:       name,
		State:      "new",
		Type:       "BaseParallelTask",
		Tasks:      make([]interface{}, 0),
		inChannel:  inchannel,
		outChannel: outchannel,
	}
}

// AddTask add parallel task with name
func (pt *BaseParallelTask) AddTask(task interface{}) {
	pt.Tasks = append(pt.Tasks, task)
}

// Execute implement Task.Execute.
func (pt *BaseParallelTask) Execute() error {
	pt.State = "inprogress"

	resultChan := make(chan error)

	for value := range pt.inChannel {

		errChan := make(chan error)
		var wg sync.WaitGroup

		for _, nt := range pt.Tasks {

			switch nttype := nt.(type) {
			case *someHumanTask:
				fmt.Println("someHumanTask - ", nttype)
				// RUN IT in separate goroutine
				wg.Add(1)
				go func(t Task) {
					if err := t.ExecuteParallel(value); err != nil {
						errChan <- err
					}
					wg.Done()
				}(nttype)

			case *extremeValueCheckTask:
				fmt.Println("extremeValueCheckTask -", nttype)
				// RUN IT in separate goroutine
				wg.Add(1)
				go func(t Task) {
					if err := t.ExecuteParallel(value); err != nil {
						errChan <- err
					}
					wg.Done()
				}(nttype)
			case *sendEmailTask:
				fmt.Println("sendEmailTask -", nttype)
				// RUN IT in separate goroutine
				wg.Add(1)
				go func(t Task) {
					if err := t.ExecuteParallel(value); err != nil {
						errChan <- err
					}
					wg.Done()
				}(nttype)
			case *sendSmsTask:
				fmt.Println("sendSmsTask -", nttype)
				// RUN IT in separate goroutine
				wg.Add(1)
				go func(t Task) {
					if err := t.ExecuteParallel(value); err != nil {
						errChan <- err
					}
					wg.Done()
				}(nttype)
			case *twitterPostTask:
				fmt.Println("twitterPostTask -", nttype)
				// RUN IT in separate goroutine
				wg.Add(1)
				go func(t Task) {
					if err := t.ExecuteParallel(value); err != nil {
						errChan <- err
					}
					wg.Done()
				}(nttype)
			case *sendToDatabase:
				fmt.Println("sendToDatabase -", nttype)
				// RUN IT in separate goroutine
				wg.Add(1)
				go func(t Task) {
					if err := t.ExecuteParallel(value); err != nil {
						errChan <- err
					}
					wg.Done()
				}(nttype)
			default:
				fmt.Println("----default-----", nttype)
				// fmt.Println(nttype)
			}

		}

		go func() {
			var result *multierror.Error
			for err := range errChan {
				result = multierror.Append(result, err)
			}
			resultChan <- result.ErrorOrNil()
		}()

		wg.Wait()
		close(errChan)

		pt.outChannel <- value
	}

	pt.State = "completed"
	return <-resultChan
}

// ExecuteParallel - default empty implementation
func (pt *BaseParallelTask) ExecuteParallel(value string) error {
	// do nothing
	return nil
}

//*********************************************
// Base Decision Task
//*********************************************
type BaseDecisionTask struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	State       string        `json:"state" bson:"state"`
	inChannel   chan string
	outChannels []chan string
}

func NewDecisionTask(name string, inchannel chan string, outChannels []chan string) *BaseDecisionTask {
	return &BaseDecisionTask{
		Name:        name,
		State:       "new",
		inChannel:   inchannel,
		outChannels: outChannels,
	}
}

// Execute implement Task.Execute.
func (t *BaseDecisionTask) Execute() error {
	t.State = "inprogress"

	for value := range t.inChannel {

		val, _ := strconv.Atoi(value)

		if val < 10 {
			t.outChannels[0] <- value
		} else if val > 10 && val <= 30 {
			t.outChannels[1] <- value
		} else if val > 30 {
			t.outChannels[2] <- value
		}

	}

	t.State = "completed"
	return nil
}

func (t *BaseDecisionTask) ExecuteParallel(value string) error {
	//nothing here
	return nil
}
