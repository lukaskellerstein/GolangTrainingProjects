package main

import (
	"./core/abstraction"
	"./core/decision"
	"./core/human"
	"./core/log"
	"./core/send"
	"./core/workflow"

	"gopkg.in/mgo.v2/bson"
)

func RunWorkflow1(name string) {

	// wfdb := workflow.GetWorkflow(&workflow.Workflow{Name: name})

	aaa := true
	if aaa {
		// NEW WORKFLOW
		go func() {
			wf := workflow.NewWorkflow(name)

			// CHANNELS ------------------

			ch1out := make(chan string)
			// ch2out := make(chan string)
			ch3out := make(chan string)

			ch4out := make(chan string)
			ch5out := make(chan string)
			ch6out := make(chan string)

			ch7out := make(chan string)

			wf.ChannelsCount = 7

			// PIPELINE ------------------

			//normal task
			wf.AddTask(&log.LogTask{
				BaseTask: abstraction.BaseTask{
					Type:            "LogTask",
					Name:            "myTask1",
					State:           "new",
					ID:              bson.NewObjectId(),
					InChannelIndex:  0,
					OutChannelIndex: 1,
					InChannel:       workflowIn,
					OutChannel:      ch1out,
				}})

			//human task
			ht := human.NewHumanTask("ht1", 2, 3, ch1out, ch3out)
			ht.UserID = "user21"
			ht.ResolvedTime = "27-10-2017 15:43:00"
			wf.AddTask(ht)

			//decision task
			chas := make([]chan string, 3)
			chas[0] = ch4out
			chas[1] = ch5out
			chas[2] = ch6out
			dt := decision.NewDecisionTask("Decision1", 3, []int{4, 5, 6}, ch3out, chas)
			wf.AddTask(dt)

			//decision's tasks
			wf.AddTask(&send.SendEmailTask{
				BaseTask: abstraction.BaseTask{
					Type:            "SendEmailTask",
					Name:            "myTask5",
					State:           "new",
					ID:              bson.NewObjectId(),
					InChannelIndex:  4,
					OutChannelIndex: 7,
					InChannel:       ch4out,
					OutChannel:      ch7out}, EmailAddress: "someuser@gmail.com"})
			wf.AddTask(&send.SendSmsTask{
				BaseTask: abstraction.BaseTask{
					Type:            "SendSmsTask",
					Name:            "myTask6",
					State:           "new",
					ID:              bson.NewObjectId(),
					InChannelIndex:  5,
					OutChannelIndex: 7,
					InChannel:       ch5out,
					OutChannel:      ch7out}, PhoneNumber: "725012034"})
			wf.AddTask(&send.SendMqttTask{
				BaseTask: abstraction.BaseTask{
					Type:            "SendMqttTask",
					Name:            "myTask7",
					State:           "new",
					ID:              bson.NewObjectId(),
					InChannelIndex:  6,
					OutChannelIndex: 7,
					InChannel:       ch6out,
					OutChannel:      ch7out}, Topic: "house/alarm"})

			//normal task
			wf.AddTask(&log.LogTask{
				BaseTask: abstraction.BaseTask{
					Type:            "LogTask",
					Name:            "myTask8",
					State:           "new",
					ID:              bson.NewObjectId(),
					InChannelIndex:  7,
					OutChannelIndex: 99,
					InChannel:       ch7out,
					OutChannel:      workflowOut,
				}})

			//---------------------------
			workflow.SaveWorkflow(wf)
			//---------------------------

			wf.Run()
		}()

	} else {
		// // EXISTING WORKFLOW
		// go func() {

		// 	// CHANNELS ------------------

		// 	ch1out := make(chan string)
		// 	ch2out := make(chan string)
		// 	ch3out := make(chan string)

		// 	ch4out := make(chan string)
		// 	ch5out := make(chan string)
		// 	ch6out := make(chan string)

		// 	ch7out := make(chan string)

		// 	channels := make([]chan string, 1) //empty like workflowin

		// 	channels = append(channels, ch1out)
		// 	channels = append(channels, ch2out)
		// 	channels = append(channels, ch3out) //ch3out := make(chan string)

		// 	channels = append(channels, ch4out) //ch4out := make(chan string)
		// 	channels = append(channels, ch5out) //ch5out := make(chan string)
		// 	channels = append(channels, ch6out) //ch6out := make(chan string)

		// 	channels = append(channels, ch7out) //ch7out := make(chan string)

		// 	// PIPELINE ------------------
		// 	wfram := &workflow.Workflow{}

		// 	for _, nt := range wfdb.Tasks {

		// 		concreteTaskVariable := nt.(bson.M)

		// 		if concreteTaskVariable["type"] == "BaseDecisionTask" {
		// 			task2 := &decision.BaseDecisionTask{}
		// 			bodyBytes, _ := json.Marshal(concreteTaskVariable)
		// 			json.Unmarshal(bodyBytes, &task2)
		// 			// fmt.Println(asdf)

		// 			//IN CHANNEL
		// 			if task2.InChannelIndex == 0 {
		// 				task2.InChannel = workflowIn
		// 			} else {
		// 				task2.InChannel = channels[task2.InChannelIndex]
		// 			}

		// 			//OUT CHANNELS
		// 			for chindx := range task2.OutChannelsIndex {
		// 				if chindx == 99 {
		// 					task2.OutChannels = append(task2.OutChannels, workflowOut)
		// 				} else {
		// 					task2.OutChannels = append(task2.OutChannels, channels[chindx])
		// 				}
		// 			}

		// 			wfram.AddTask(task2)
		// 		} else {

		// 			baseTaskVariable := nt.(bson.M)["basetask"]
		// 			//BaseTask
		// 			asdf := abstraction.BaseTask{}
		// 			bodyBytes, _ := json.Marshal(baseTaskVariable)
		// 			json.Unmarshal(bodyBytes, &asdf)
		// 			// fmt.Println(asdf)

		// 			//COCNRETE TASK
		// 			if asdf.Type == "LogTask" {
		// 				asdf2 := &log.LogTask{BaseTask: asdf}
		// 				bodyBytes, _ := json.Marshal(concreteTaskVariable)
		// 				json.Unmarshal(bodyBytes, &asdf2)

		// 				// fmt.Println(asdf2)
		// 				// fmt.Println(asdf2.ID)
		// 				// fmt.Println(asdf2.Name)
		// 				// fmt.Println(asdf2.MinValue)
		// 				// fmt.Println(asdf2.MaxValue)

		// 				//IN CHANNEL
		// 				if asdf2.InChannelIndex == 0 {
		// 					asdf2.InChannel = workflowIn
		// 				} else {
		// 					asdf2.InChannel = channels[asdf2.InChannelIndex]
		// 				}

		// 				//OUT CHANNEL
		// 				if asdf2.OutChannelIndex == 99 {
		// 					asdf2.OutChannel = workflowOut
		// 				} else {
		// 					asdf2.OutChannel = channels[asdf2.OutChannelIndex]
		// 				}

		// 				wfram.AddTask(asdf2)
		// 			} else if asdf.Type == "BaseHumanTask" {
		// 				asdf2 := &human.BaseHumanTask{BaseTask: asdf}
		// 				bodyBytes, _ := json.Marshal(concreteTaskVariable)
		// 				json.Unmarshal(bodyBytes, &asdf2)

		// 				// fmt.Println(asdf2)
		// 				// fmt.Println(asdf2.ID)
		// 				// fmt.Println(asdf2.Name)
		// 				// fmt.Println(asdf2.UserID)
		// 				// fmt.Println(asdf2.ResolvedTime)

		// 				//IN CHANNEL
		// 				if asdf2.InChannelIndex == 0 {
		// 					asdf2.InChannel = workflowIn
		// 				} else {
		// 					asdf2.InChannel = channels[asdf2.InChannelIndex]
		// 				}

		// 				//OUT CHANNEL
		// 				if asdf2.OutChannelIndex == 99 {
		// 					asdf2.OutChannel = workflowOut
		// 				} else {
		// 					asdf2.OutChannel = channels[asdf2.OutChannelIndex]
		// 				}

		// 				wfram.AddTask(asdf2)
		// 			} else if asdf.Type == "SendEmailTask" {
		// 				asdf2 := &send.SendEmailTask{BaseTask: asdf}
		// 				bodyBytes, _ := json.Marshal(concreteTaskVariable)
		// 				json.Unmarshal(bodyBytes, &asdf2)

		// 				// fmt.Println(asdf2)
		// 				// fmt.Println(asdf2.ID)
		// 				// fmt.Println(asdf2.Name)
		// 				// fmt.Println(asdf2.EmailAddress)

		// 				//IN CHANNEL
		// 				if asdf2.InChannelIndex == 0 {
		// 					asdf2.InChannel = workflowIn
		// 				} else {
		// 					asdf2.InChannel = channels[asdf2.InChannelIndex]
		// 				}

		// 				//OUT CHANNEL
		// 				if asdf2.OutChannelIndex == 99 {
		// 					asdf2.OutChannel = workflowOut
		// 				} else {
		// 					asdf2.OutChannel = channels[asdf2.OutChannelIndex]
		// 				}

		// 				wfram.AddTask(asdf2)
		// 			} else if asdf.Type == "SendSmsTask" {
		// 				asdf2 := &send.SendSmsTask{BaseTask: asdf}
		// 				bodyBytes, _ := json.Marshal(concreteTaskVariable)
		// 				json.Unmarshal(bodyBytes, &asdf2)

		// 				// fmt.Println(asdf2)
		// 				// fmt.Println(asdf2.ID)
		// 				// fmt.Println(asdf2.Name)

		// 				//IN CHANNEL
		// 				if asdf2.InChannelIndex == 0 {
		// 					asdf2.InChannel = workflowIn
		// 				} else {
		// 					asdf2.InChannel = channels[asdf2.InChannelIndex]
		// 				}

		// 				//OUT CHANNEL
		// 				if asdf2.OutChannelIndex == 99 {
		// 					asdf2.OutChannel = workflowOut
		// 				} else {
		// 					asdf2.OutChannel = channels[asdf2.OutChannelIndex]
		// 				}

		// 				wfram.AddTask(asdf2)
		// 			} else if asdf.Type == "SendMqttTask" {
		// 				asdf2 := &send.SendMqttTask{BaseTask: asdf}
		// 				bodyBytes, _ := json.Marshal(concreteTaskVariable)
		// 				json.Unmarshal(bodyBytes, &asdf2)

		// 				// fmt.Println(asdf2)
		// 				// fmt.Println(asdf2.ID)
		// 				// fmt.Println(asdf2.Name)

		// 				//IN CHANNEL
		// 				if asdf2.InChannelIndex == 0 {
		// 					asdf2.InChannel = workflowIn
		// 				} else {
		// 					asdf2.InChannel = channels[asdf2.InChannelIndex]
		// 				}

		// 				//OUT CHANNEL
		// 				if asdf2.OutChannelIndex == 99 {
		// 					asdf2.OutChannel = workflowOut
		// 				} else {
		// 					asdf2.OutChannel = channels[asdf2.OutChannelIndex]
		// 				}

		// 				wfram.AddTask(asdf2)
		// 			} else if asdf.Type == "SendRpcTask" {
		// 				asdf2 := &send.SendRpcTask{BaseTask: asdf}
		// 				bodyBytes, _ := json.Marshal(concreteTaskVariable)
		// 				json.Unmarshal(bodyBytes, &asdf2)

		// 				// fmt.Println(asdf2)
		// 				// fmt.Println(asdf2.ID)
		// 				// fmt.Println(asdf2.Name)
		// 				// fmt.Println(asdf2.DatabaseName)

		// 				//IN CHANNEL
		// 				if asdf2.InChannelIndex == 0 {
		// 					asdf2.InChannel = workflowIn
		// 				} else {
		// 					asdf2.InChannel = channels[asdf2.InChannelIndex]
		// 				}

		// 				//OUT CHANNEL
		// 				if asdf2.OutChannelIndex == 99 {
		// 					asdf2.OutChannel = workflowOut
		// 				} else {
		// 					asdf2.OutChannel = channels[asdf2.OutChannelIndex]
		// 				}

		// 				wfram.AddTask(asdf2)
		// 			}

		// 		}

		// 	}
		// 	//_________________________________________________

		// 	wfram.Run()

		// }()
	}
}
