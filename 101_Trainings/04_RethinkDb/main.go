package main

import (
	"fmt"
	"log"
	
	r "gopkg.in/gorethink/gorethink.v3"
)

type person struct{
	id string
	name string
	surname string
}

func (p person) getName(){
	fmt.Println(p.name)
}



func main(){
	fmt.Println("START")
	
	//Object ----------------------------------------
	newOne := person{"afda0sfdsr33", "Lukas", "Kellerstein"}
	newOne.getName()
	// --------------------------------------------------


	//RethinkDB ----------------------------------------
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
		Database: "test",
	})
	if err != nil {
		log.Fatalln(err)
	}

	res, err := r.Expr("Hello World").Run(session)
	if err != nil {
		log.Fatalln(err)
	}

	var response string
	err = res.One(&response)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(response)

	// --------------------------------------------------

	fmt.Println("END")
}