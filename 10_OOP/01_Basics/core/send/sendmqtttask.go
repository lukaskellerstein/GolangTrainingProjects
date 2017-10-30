package send

import (
	"fmt"

	"../abstraction"
)

//**********************************
//TASK - Send mqtt
//**********************************
type SendMqttTask struct {
	abstraction.BaseTask
	Topic string `json:"topic" bson:"topic"`
}

func (t *SendMqttTask) Execute() error {
	fmt.Println("SendMqttTask execute")

	for value := range t.InChannel {
		t.State = "inprogress"
		//*****************
		// DOING SOMETHING
		fmt.Println("SendMqttTask value - " + value)
		//*****************
		t.OutChannel <- value
	}

	t.State = "completed"
	return nil
}

func (t *SendMqttTask) ExecuteParallel(value string) error {
	t.State = "inprogress"
	//*****************
	// DOING SOMETHING
	fmt.Println("SendMqttTask parallel value - " + value)
	//*****************
	t.State = "completed"
	return nil
}
