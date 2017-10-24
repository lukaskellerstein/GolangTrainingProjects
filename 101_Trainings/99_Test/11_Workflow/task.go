package main

import (
	"gopkg.in/mgo.v2/bson"
)

// Task represents task interface of workflow.
type Task interface {
	// SetID(id bson.ObjectId)
	// GetID() bson.ObjectId
	// SetName(name string)
	// GetName() string
	// SetState(state string)
	// GetState() string
	Execute() error
}

type BaseTask struct {
	ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	State      string        `json:"state" bson:"state"`
	Value      string        `json:"value" bson:"value"`
	inChannel  chan string
	outChannel chan string
}

func (t *BaseTask) Execute() error {
	// do nothing
	return nil
}
