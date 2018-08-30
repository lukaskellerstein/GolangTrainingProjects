package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type Task interface {
	Execute()
}

type BaseTask struct {
	ChannelIn    chan string
	ChannelOut   chan string
	ChannelClose chan string
}

type ConcreteTask1 struct {
	BaseTask
	Name string
}

func (t *ConcreteTask1) Execute() {

	go func() {
		<-t.ChannelClose
		fmt.Println("closing ConcreteTask1")
		close(t.ChannelOut)
	}()

	for item := range t.ChannelIn {
		item += "1"
		fmt.Println(item)
		t.ChannelOut <- item
	}
}

type ConcreteTask2 struct {
	BaseTask
	Description string
}

func (t *ConcreteTask2) Execute() {

	go func() {
		<-t.ChannelClose
		fmt.Println("closing ConcreteTask2")
		close(t.ChannelOut)
	}()

	for item := range t.ChannelIn {
		item += "2"
		fmt.Println(item)
		t.ChannelOut <- item
	}
}

type ConcreteTask3 struct {
	BaseTask
	Type string
}

func (t *ConcreteTask3) Execute() {

	go func() {
		<-t.ChannelClose
		fmt.Println("closing ConcreteTask3")
		close(t.ChannelOut)
	}()

	for item := range t.ChannelIn {
		item += "3"
		fmt.Println(item)
		t.ChannelOut <- item
	}
}

func main() {

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	main_channelClose := make(chan string)

	ct1_channelIn := make(chan string)
	ct1_channelClose := make(chan string)

	ct2_channelIn := make(chan string)
	ct2_channelClose := make(chan string)

	ct3_channelIn := make(chan string)
	ct3_channelOut := make(chan string)
	ct3_channelClose := make(chan string)

	ct1 := ConcreteTask1{
		BaseTask: BaseTask{
			ChannelIn:    ct1_channelIn,
			ChannelOut:   ct2_channelIn,
			ChannelClose: ct1_channelClose,
		},
		Name: "someName",
	}

	ct2 := ConcreteTask2{
		BaseTask: BaseTask{
			ChannelIn:    ct2_channelIn,
			ChannelOut:   ct3_channelIn,
			ChannelClose: ct2_channelClose,
		},
		Description: "someDesc",
	}

	ct3 := ConcreteTask3{
		BaseTask: BaseTask{
			ChannelIn:    ct3_channelIn,
			ChannelOut:   ct3_channelOut,
			ChannelClose: ct3_channelClose,
		},
		Type: "someType",
	}

	go ct1.Execute()
	go ct2.Execute()
	go ct3.Execute()

	go func() {
		for item := range ct3.ChannelOut {
			fmt.Println(item)
		}
	}()

	go func() {
	loop:
		for {
			time.Sleep(time.Duration(1) * time.Second)
			randomNumberFloat := rand.Float64() * 1000

			select {
			case <-main_channelClose:
				break loop // has to be named, because "break" applies to the select otherwise
			default:
				//do nothing
			}

			ct1.ChannelIn <- strconv.FormatFloat(randomNumberFloat, 'E', -1, 64)
		}
	}()

	go func() {
		time.Sleep(time.Duration(15) * time.Second)

		main_channelClose <- "close sending"
		ct1.ChannelClose <- "close it"
		ct2.ChannelClose <- "close it"
		ct3.ChannelClose <- "close it"
	}()

	<-sigc
}
