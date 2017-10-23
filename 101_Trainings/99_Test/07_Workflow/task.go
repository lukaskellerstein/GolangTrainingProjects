package main

import (
	"strconv"

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

type DecisionTask struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	State       string        `json:"state" bson:"state"`
	inChannel   chan string
	outChannels [3]chan string
}

func NewDecisionTask(name string, inchannel chan string, outChannels [3]chan string) *DecisionTask {
	return &DecisionTask{
		Name:        name,
		State:       "new",
		inChannel:   inchannel,
		outChannels: outChannels,
	}
}

func (t *DecisionTask) SetID(id bson.ObjectId) {
	t.ID = id
}

func (t *DecisionTask) GetID() bson.ObjectId {
	return t.ID
}

func (t *DecisionTask) SetName(name string) {
	t.Name = name
}

func (t *DecisionTask) GetName() string {
	return t.Name
}

func (t *DecisionTask) SetState(state string) {
	t.State = state
}

func (t *DecisionTask) GetState() string {
	return t.State
}

// Execute implement Task.Execute.
func (t *DecisionTask) Execute() error {
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

func (t *DecisionTask) ExecuteParallel(value string) error {
	//nothing here
	return nil
}
