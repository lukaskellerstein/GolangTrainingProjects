package main

import (
	"fmt"
	"log"

	r "gopkg.in/gorethink/gorethink.v3"
)

func main() {
	fmt.Println("START")

	//RethinkDB ----------------------------------------
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "test",
	})
	if err != nil {
		log.Fatalln(err)
	}

	//Recreate tables
	err = r.DB("test").TableDrop("posts").Exec(session)
	err = r.DB("test").TableCreate("posts").Exec(session)

	//Insert into RethinkDB
	res, err := r.DB("test").Table("posts").Insert(map[string]string{
		"id":      "1",
		"title":   "Lorem ipsum",
		"content": "Dolor sit amet",
	}).Run(session)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)

	// --------------------------------------------------

	fmt.Println("END")
}
