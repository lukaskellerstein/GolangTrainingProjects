package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	//Encode
	p1 := person{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1)

	//Decode
	var p2 person
	rdr := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.NewDecoder(rdr).Decode(&p2)

	fmt.Println(p2.First)
	fmt.Println(p2.Last)
	fmt.Println(p2.Age)
}
