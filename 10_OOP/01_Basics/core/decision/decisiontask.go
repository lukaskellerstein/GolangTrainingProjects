package decision

import (
	"fmt"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

//*********************************************
// Base Decision Task
//*********************************************
type BaseDecisionTask struct {
	ID               bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name             string        `json:"name" bson:"name"`
	State            string        `json:"state" bson:"state"`
	Type             string        `json:"type" bson:"type"`
	InChannelIndex   int           `json:"inchannelindex" bson:"inchannelindex"`
	OutChannelsIndex []int         `json:"outchannelsindex" bson:"outchannelsindex"`
	InChannel        chan string   `json:"-" bson:"-"`
	OutChannels      []chan string `json:"-" bson:"-"`
}

func NewDecisionTask(name string, inchannelindex int, outchannelsindex []int, inchannel chan string, outChannels []chan string) *BaseDecisionTask {
	return &BaseDecisionTask{
		Name:             name,
		State:            "new",
		Type:             "BaseDecisionTask",
		InChannelIndex:   inchannelindex,
		OutChannelsIndex: outchannelsindex,
		InChannel:        inchannel,
		OutChannels:      outChannels,
	}
}

// Execute implement Task.Execute.
func (t *BaseDecisionTask) Execute() error {
	fmt.Println("BaseDecisionTask execute")

	t.State = "inprogress"

	for value := range t.InChannel {

		val, _ := strconv.Atoi(value)

		if val < 10 {
			t.OutChannels[0] <- value
		} else if val > 10 && val <= 30 {
			t.OutChannels[1] <- value
		} else if val > 30 {
			t.OutChannels[2] <- value
		}

	}

	t.State = "completed"
	return nil
}

func (t *BaseDecisionTask) ExecuteParallel(value string) error {
	//nothing here
	return nil
}
