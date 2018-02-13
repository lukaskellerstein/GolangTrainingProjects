package main

import (
	"testing"

	mgo "gopkg.in/mgo.v2"
)

func Benchmark_saveToDB_One(b *testing.B) {
	var tempObject = TestObject{
		FirstName: "test1",
		LastName:  "test2",
	}

	for index := 0; index < b.N; index++ {
		saveToDB_One(tempObject)
	}
}

func Benchmark_saveToDB_Two(b *testing.B) {
	var tempObject = TestObject{
		FirstName: "test1",
		LastName:  "test2",
	}

	for index := 0; index < b.N; index++ {
		saveToDB_Two(&tempObject)
	}
}

func Benchmark_saveToDB_Three(b *testing.B) {

	session, err = mgo.Dial(mongodburl)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var tempObject = TestObject{
		FirstName: "test1",
		LastName:  "test2",
	}

	for index := 0; index < b.N; index++ {
		saveToDB_Three(tempObject)
	}
}

func Benchmark_saveToDB_Four(b *testing.B) {

	session, err = mgo.Dial(mongodburl)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var tempObject = TestObject{
		FirstName: "test1",
		LastName:  "test2",
	}

	for index := 0; index < b.N; index++ {
		saveToDB_Four(&tempObject)
	}
}
