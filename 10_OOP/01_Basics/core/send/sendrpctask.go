package send

import (
	"fmt"

	abs "../abstraction"
)

//**********************************
//TASK - Send gRPC
//**********************************
type SendRpcTask struct {
	abs.BaseTask
	address string `json:"address" bson:"address"`
}

func (t *SendRpcTask) Execute() error {
	fmt.Println("SendRpcTask execute")

	for value := range t.InChannel {
		t.State = "inprogress"
		//*****************
		// DOING SOMETHING
		fmt.Println("SendRpcTask value - " + value)
		//*****************
		t.OutChannel <- value
	}

	t.State = "completed"
	return nil
}

func (t *SendRpcTask) ExecuteParallel(value string) error {
	t.State = "inprogress"
	//*****************
	// DOING SOMETHING
	fmt.Println("SendRpcTask parallel value - " + value)
	//*****************
	t.State = "completed"
	return nil
}
