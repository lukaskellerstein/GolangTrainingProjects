package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(db AppDatabase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home page - "+db.GetText())
	})
}

func main() {
	db := NewMongoDatabase()

	http.Handle("/home", HomeHandler(db))
	http.ListenAndServe(":8086", nil)
}
