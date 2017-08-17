package main

import "fmt"
import "github.com/lukaskellerstein/GolangTrainingProjects/01_Packages1/somepackage1"

func main() {
	aaa := ""

	bbb := "asdfasfasdf"

	fmt.Println(aaa)
	fmt.Println(bbb)

	ccc := somepackage1.GetData()
	fmt.Println(ccc)
}
