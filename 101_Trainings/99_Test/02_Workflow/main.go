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

	buf.WriteString("3.c<ParallelTask>(c1<summaryTask>, c2<summaryTask>)")
	pt := NewParallelTask()
	pt.AddTask("c1", &myTask1{})
	pt.AddTask("c2", &myTask1{})
	wf.AddTask("c", pt)
	buf.WriteString(" -> ")

	buf.WriteString("4.d<Workflow>(1.da<summaryTask> -> 2.db<summaryTask> -> 3.dc<summaryTask>)")
	wf2 := NewWorkflow()
	wf2.AddTask("da", &myTask1{})
	wf2.AddTask("db", &myTask1{})
	wf2.AddTask("dc", &myTask1{})
	wf.AddTask("d", wf2)

	expect := buf.String()
	if wf.Summary() != expect {
		fmt.Printf("workflow summary \ngot:   %v\nexpect:%v", wf.Summary(), expect)
	}

	fmt.Println("END")
}
