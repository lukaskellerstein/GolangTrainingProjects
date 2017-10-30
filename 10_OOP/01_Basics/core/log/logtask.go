package log

import (
	"fmt"

	"../abstraction"
)

//**********************************
//TASK - log
//**********************************
type LogTask struct {
	abstraction.BaseTask
}

func (t *LogTask) Execute() error {

	fmt.Println("LogTask execute")

	for value := range t.InChannel {

		fmt.Println("LogTask value - " + value)

		t.OutChannel <- value
	}

	//SEM SE TO NEDOSTANE ??? vim proc
	fmt.Println("LogTask - COMPLETE")

	//wg.Done()
	t.State = "completed"
	return nil
}

func (t *LogTask) ExecuteParallel(value string) error {
	t.State = "inprogress"
	// t.Value = value
	fmt.Println("LogTask parallel value - " + value)
	t.State = "completed"
	return nil
}
