package main

import (
	"../package1"
	"../package2"
	"fmt"
)

func main() {

	testData1 := package1.GetData()
	testData2 := package2.GetData()

	fmt.Println(testData1)
	fmt.Println(testData2)

	fmt.Println(package1.SomeExportedVar)
	fmt.Println(package2.SomeExportedVar)
}
