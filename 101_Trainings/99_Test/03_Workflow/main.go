package main

import (
	"bytes"
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

//**********************************
//TASK 3
//**********************************
type myTask3 struct {
	name string
}

func (t *myTask3) Execute() error {
	fmt.Println("task3 - " + t.name)
	return nil
}

//**********************************
//TASK 4
//**********************************
type myTask4 struct {
	name string
}

func (t *myTask4) Execute() error {
	fmt.Println("task4 - " + t.name)
	return nil
}

func main() {
	fmt.Println("START")

	buf := bytes.NewBufferString("")

	wf := NewWorkflow()

	buf.WriteString("1.a<summaryTask>")
	wf.AddTask("a", &myTask1{})
	buf.WriteString(" -> ")

	buf.WriteString("2.b<summaryTask>")
	wf.AddTask("b", &myTask1{})
	buf.WriteString(" -> ")

	buf.WriteString("3.c<ParallelTask>(c1<myTask2>, c2<myTask3>, c2<myTask4>)")
	pt := NewParallelTask()
	pt.AddTask("c1", &myTask2{})
	pt.AddTask("c2", &myTask3{})
	pt.AddTask("c3", &myTask4{})
	wf.AddTask("c", pt)
	buf.WriteString(" -> ")

	expect := buf.String()
	if wf.Summary() != expect {
		fmt.Printf("workflow summary \ngot:   %v\nexpect:%v", wf.Summary(), expect)
	}

	fmt.Println("END")
}
