package main

import (
	"net/http"
)

//BAD WAY - cannot be tested
// func homeHandleFunc(w http.ResponseWriter, r *http.Request) {
// w.Header().Set("Content-Type", "text/html; charset=utf-8")
// w.Write([]byte("<div>Home page</div>"))
// }

//BETTER WAY - can be tested
func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<div>Home page</div>"))
	})
}

func main() {

	//BAD WAY - cannot be tested
	//http.HandleFunc("/home", homeHandleFunc)

	//BETTER WAY - can be tested
	http.Handle("/home", homeHandler())

	http.ListenAndServe(":8085", nil)
}
