package main

import (
	"fmt"
	"log"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

const (
	MyDB     = "customdatabase"
	username = "test"
	password = "test"
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
	//CREATE DATABASE
	//********************************
	ress, err := queryDB(c, fmt.Sprintf("CREATE DATABASE %s", MyDB))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ress)

	//********************************
	// INSERT DATA
	//********************************
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	tags := map[string]string{"someMetric": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}

	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
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
