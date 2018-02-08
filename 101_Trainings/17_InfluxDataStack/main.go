package main

import (
	"fmt"
	"log"

	client "github.com/influxdata/influxdb/client/v2"
)

const (
	MyDB     = "telegraf.autogen"
	username = "bubba"
	password = "bumblebeetuna"
)

func main() {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
		// Username: username,
		// Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}

	//********************************
	// SELECT DATA
	//********************************
	q := fmt.Sprintf("SELECT count(%s) FROM %s", "usage_system", "telegraf.autogen.cpu")
	res, err := queryDB(c, q)
	if err != nil {
		log.Fatal(err)
	}
	count := res[0].Series[0].Values[0][1]
	log.Printf("Found a total of %v records\n", count)
}

//********************************
// HELPER METHOD
//********************************
// queryDB convenience function to query the database
func queryDB(clnt client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: MyDB,
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}
