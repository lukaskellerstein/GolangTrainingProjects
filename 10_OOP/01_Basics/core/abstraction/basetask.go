package abstraction

import "gopkg.in/mgo.v2/bson"

type BaseTask struct {
	ID    bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name  string        `json:"name" bson:"name"`
	Type  string        `json:"type" bson:"type"`
	State string        `json:"state" bson:"state"`
	// Value           string        `json:"value" bson:"value"`
	InChannelIndex  int         `json:"inchannelindex" bson:"inchannelindex"`
	OutChannelIndex int         `json:"outchannelindex" bson:"outchannelindex"`
	InChannel       chan string `json:"-" bson:"-"`
	OutChannel      chan string `json:"-" bson:"-"`
}

// // default empty implementation
// func (t *BaseTask) Execute() error {
// 	// do nothing
// 	return nil
// }
