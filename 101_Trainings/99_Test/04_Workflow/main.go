package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func main() {
	fmt.Println("START")

	wf := NewWorkflow("Workflow1")
	wf.AddTask(&extremeValueCheckTask{})

	// pt := NewParallelTask()
	// pt.AddTask(&sendEmailTask{})
	// pt.AddTask(&sendSmsTask{})
	// pt.AddTask(&twitterPostTask{})
	// wf.AddTask(pt)

	// wf.AddTask(&extremeValueCheckTask{})

	wf.Run()

	//save to the mongodb
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
