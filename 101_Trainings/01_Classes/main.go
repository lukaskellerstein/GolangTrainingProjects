package main

import "fmt"


type person struct{
	id string
	name string
	surname string
}

func (p person) getName(){
	fmt.Println(p.name)
}


func main(){

	newOne := person{"afda0sfdsr33", "Lukas", "Kellerstein"}
	newOne.getName()

	fmt.Println("sometext")
}