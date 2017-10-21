package main

import (
	"fmt"
)

//**********************************
//TASK 1
//**********************************
type myTask1 struct {
	name string
}

func (t *myTask1) Execute() error {
	fmt.Println("task1 - " + t.name)
	return nil
}

//**********************************
//TASK 2
//**********************************
type myTask2 struct {
	name string
}

func (t *myTask2) Execute() error {
	fmt.Println("task2 - " + t.name)
	return nil
}

func main() {
	fmt.Println("START")

	wf := NewWorkflow()
	wf.AddTask("a1", &myTask1{name: "a1"})
	wf.AddTask("a2", &myTask1{name: "a2"})
	wf.AddTask("a3", &myTask1{name: "a3"})
	wf.AddTask("b1", &myTask2{name: "b1"})
	wf.AddTask("b2", &myTask2{name: "b2"})
	wf.AddTask("b3", &myTask2{name: "b3"})
	wf.AddTask("b4", &myTask2{name: "b4"})
	err := wf.Run()

	if err != nil {
		panic(err)
	}

	fmt.Println("END")
}
