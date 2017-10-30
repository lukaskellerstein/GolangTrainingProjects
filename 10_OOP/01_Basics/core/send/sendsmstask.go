package send

import (
	"fmt"

	abs "../abstraction"
)

//**********************************
//TASK - Send SMS
//**********************************
type SendSmsTask struct {
	abs.BaseTask
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber"`
}

func (t *SendSmsTask) Execute() error {
	fmt.Println("SendSmsTask execute")

	for value := range t.InChannel {
		t.State = "inprogress"
		//*****************
		// DOING SOMETHING
		fmt.Println("SendSmsTask value - " + value)
		//*****************
		t.OutChannel <- value
	}

	//SEM SE TO NEDOSTANE ???
	fmt.Println("SendSmsTask - COMPLETE")

	t.State = "completed"
	return nil
}

func (t *SendSmsTask) ExecuteParallel(value string) error {
	t.State = "inprogress"
	//*****************
	// DOING SOMETHING
	fmt.Println("SendSmsTask parallel value - " + value)
	//*****************
	t.State = "completed"
	return nil
}
