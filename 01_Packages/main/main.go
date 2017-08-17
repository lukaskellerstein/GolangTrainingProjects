package main

import "fmt"
import "../somepackage1"
import "../somepackage2"

func main() {
	aaa := ""

	bbb := "asdfasfasdf"

	fmt.Println(aaa)
	fmt.Println(bbb)

	ccc := somepackage1.GetData()
	fmt.Println(ccc)

	ddd := somepackage2.GetData()
	fmt.Println(ddd)
}
