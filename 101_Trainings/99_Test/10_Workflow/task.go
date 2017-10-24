package main

import (
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
