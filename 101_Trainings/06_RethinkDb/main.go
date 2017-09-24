package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	r "gopkg.in/gorethink/gorethink.v3"
)

//Person - class for Person object
type Person struct {
	ID    string `gorethink:"id,omitempty"`
	Name  string `gorethink:"name"`
	Score int    `gorethink:"score"`
}

func (p Person) getName() {
	fmt.Println(p.Name)
}

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
	err = r.DB("test").TableDrop("users").Exec(session)
	err = r.DB("test").TableCreate("users").Exec(session)

	//Insert into RethinkDB
	for i := 0; i < 100; i++ {
		player := new(Person)
		player.ID = strconv.Itoa(i)
		player.Name = fmt.Sprintf("Player %d", i)
		player.Score = rand.Intn(100)

		_, err := r.DB("test").Table("users").Insert(player).RunWrite(session)

		if err != nil {
			log.Fatal(err)
		}
	}

	// --------------------------------------------------

	fmt.Println("END")
}
