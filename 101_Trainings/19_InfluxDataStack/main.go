package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

const (
	MyDB     = "customdatabase2"
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
	writePoints(c)
}

func writePoints(clnt client.Client) {
	sampleSize := 1000

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "us",
	})
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < sampleSize; i++ {
		regions := []string{"us-west1", "us-west2", "us-west3", "us-east1"}
		tags := map[string]string{
			"cpu":    "cpu-total",
			"host":   fmt.Sprintf("host%d", rand.Intn(1000)),
			"region": regions[rand.Intn(len(regions))],
		}

		idle := rand.Float64() * 100.0
		fields := map[string]interface{}{
			"idle": idle,
			"busy": 100.0 - idle,
		}

		pt, err := client.NewPoint(
			"cpu_usage",
			tags,
			fields,
			time.Now(),
		)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}

	if err := clnt.Write(bp); err != nil {
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
