package implementation

import (
	"fmt"
	"test/cellarstone1/pkg2/abstraction"
)

//**********************************
//TASK - Is there some extreme value ?
//**********************************
type ExtremeValueCheckTask struct {
	abstraction.BaseTask
	MinValue string `json:"minValue" bson:"minValue"`
	MaxValue string `json:"maxValue" bson:"maxValue"`
}

func (t *ExtremeValueCheckTask) Execute() error {

	fmt.Println("ExtremeValueCheckTask execute")

	for value := range t.InChannel {

		if t.State == "new" {
			t.State = "inprogress"
		}

		t.State = "inprogress"
		t.Value = value

		fmt.Println("ExtremeValueCheckTask value - " + t.Value)

		t.OutChannel <- value
	}

	//SEM SE TO NEDOSTANE ??? vim proc
	fmt.Println("ExtremeValueCheckTask - COMPLETE")

	//wg.Done()
	t.State = "completed"
	return nil
}

func (t *ExtremeValueCheckTask) ExecuteParallel(value string) error {
	t.State = "inprogress"
	t.Value = value
	fmt.Println("ExtremeValueCheckTask parallel value - " + t.Value)
	t.State = "completed"
	return nil
}
