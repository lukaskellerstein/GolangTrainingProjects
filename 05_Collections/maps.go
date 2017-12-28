package main

import (
	"fmt"
)

func main2() {
	// ************************
	// Maps(= Dictionary)
	// ************************

	//Full declaration
	var map1 = make(map[string]string)
	map1["ID-123213"] = "Item1"
	map1["ID-453453"] = "Item2"
	map1["ID-976484"] = "Item3"

	//Shorthand
	map2 := map[string]string{
		"ID-123213": "Item1",
		"ID-453453": "Item2",
		"ID-976484": "Item3",
	}
	fmt.Println(map2)

	//Ordinary table
	map3 := map[int]string{
		0: "Item1",
		1: "Item2",
		2: "Item3",
		3: "Item4",
	}
	fmt.Println(map3)

	//Get lenght of map
	fmt.Println(len(map2))

	//Add value
	map2["ID-somenewid"] = "somenewvalue"

	//Delete value
	delete(map2, "ID-somenewid")

	//foreach
	for key, val := range map3 {
		fmt.Println(key, " - ", val)
	}

	fmt.Println("END")
}
