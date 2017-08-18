package main

import "fmt"
import "encoding/json"

type person struct {
	First string
	Last  string `json:"Surename"` // this property will have different name in JSON
	// this property will not be exported - because it has small first letter
	age int
}

func main() {

	// create some person
	p1 := person{"Lukas", "Kellerstein", 25}

	// transfer that person to JSON
	bs, _ := json.Marshal(p1)

	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))

	// transfer that bytes back to person
	var p2 person
	json.Unmarshal(bs, &p2)

	fmt.Println("--------------")
	fmt.Println(p2.First)
	fmt.Println(p2.Last)
	fmt.Println(p2.age)

}
