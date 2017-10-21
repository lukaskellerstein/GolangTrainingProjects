package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func main() {
	fmt.Println("START")

	//*******************************
	// WORKFLOW
	//*******************************
	wf := NewWorkflow("Workflow1")

	task1 := extremeValueCheckTask{Name: "extremeValueCheckTask"}
	wf.AddTask(&task1)

	pt := NewParallelTask("ParallelTask1")
	pt.AddTask(&sendEmailTask{Name: "sendEmailTask"})
	pt.AddTask(&sendSmsTask{Name: "sendSmsTask"})
	pt.AddTask(&twitterPostTask{Name: "twitterPostTask"})
	wf.AddTask(pt)

	wf.AddTask(&sendToDatabase{Name: "sendToDatabase"})

	wf.Run()

	//*******************************
	//save to the mongodb
	//*******************************
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//SELECT TABLE
	workflowTable := session.DB("test").C("workflows")

	//INSERT
	err = workflowTable.Insert(&wf)

	if err != nil {
		panic(err)
	}

	fmt.Println("END")
}
