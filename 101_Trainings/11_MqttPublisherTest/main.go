package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("START")

	wg.Add(50)

	go senzorDHT22("s2310")
	go senzorDHT22("s2311")
	go senzorDHT22("s2312")
	go senzorDHT22("s2313")
	go senzorDHT22("s2314")
	go senzorDHT22("s2315")
	go senzorDHT22("s2316")
	go senzorDHT22("s2317")
	go senzorDHT22("s2318")
	go senzorDHT22("s2319")

	go senzorDHT22("s2320")
	go senzorDHT22("s2321")
	go senzorDHT22("s2322")
	go senzorDHT22("s2323")
	go senzorDHT22("s2324")
	go senzorDHT22("s2325")
	go senzorDHT22("s2326")
	go senzorDHT22("s2327")
	go senzorDHT22("s2328")
	go senzorDHT22("s2329")

	go senzorDHT22("s2330")
	go senzorDHT22("s2331")
	go senzorDHT22("s2332")
	go senzorDHT22("s2333")
	go senzorDHT22("s2334")
	go senzorDHT22("s2335")
	go senzorDHT22("s2336")
	go senzorDHT22("s2337")
	go senzorDHT22("s2338")
	go senzorDHT22("s2339")

	go senzorDHT22("s2340")
	go senzorDHT22("s2341")
	go senzorDHT22("s2342")
	go senzorDHT22("s2343")
	go senzorDHT22("s2344")
	go senzorDHT22("s2345")
	go senzorDHT22("s2346")
	go senzorDHT22("s2347")
	go senzorDHT22("s2348")
	go senzorDHT22("s2349")

	go senzorDHT22("s2350")
	go senzorDHT22("s2351")
	go senzorDHT22("s2352")
	go senzorDHT22("s2353")
	go senzorDHT22("s2354")
	go senzorDHT22("s2355")
	go senzorDHT22("s2356")
	go senzorDHT22("s2357")
	go senzorDHT22("s2358")
	go senzorDHT22("s2359")

	wg.Wait()

	fmt.Println("END")
}

//SIMULATE ONE SENZOR
func senzorDHT22(name string) {
	wg.Add(2)

	go doItMultipleTimes(1000, name+"/temperature")
	go doItMultipleTimes(1000, name+"/humidity")

	wg.Wait()
}

//ONE CLIENT SEND MULTIPLE TIMES
func doItMultipleTimes(count int, topic string) {

	// Create an MQTT Client.
	cli := client.New(&client.Options{
		// Define the processing of the error handler.
		ErrorHandler: func(err error) {
			fmt.Println("1", err)
		},
	})

	// Terminate the Client.
	defer cli.Terminate()

	// Connect to the MQTT Server.
	err := cli.Connect(&client.ConnectOptions{
		Network:  "tcp",
		Address:  "127.0.0.1:1883",
		ClientID: []byte("example-client-" + topic),
	})
	if err != nil {
		fmt.Println("2", err)
		panic(err)
	}

	for index := 0; index <= count; index++ {
		randomNumber := rand.Intn(50)
		sendMessageToMqttBroker(cli, topic, strconv.Itoa(randomNumber))
		time.Sleep(time.Second)
	}
}

//SIMPLE SEND MESSAGE
func sendMessageToMqttBroker(cli *client.Client, topic string, value string) {

	//*****************************
	// Publish a message.
	//*****************************
	err := cli.Publish(&client.PublishOptions{
		QoS:       mqtt.QoS0,
		TopicName: []byte(topic),
		Message:   []byte(value),
	})
	if err != nil {
		fmt.Println("3", err)
		panic(err)
	}

	fmt.Println(string(topic), string(value))
}
