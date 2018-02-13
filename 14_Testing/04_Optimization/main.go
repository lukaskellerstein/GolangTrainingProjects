package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type add struct {
	Sum int
}

var templates = template.Must(template.ParseFiles("template.gohtml"))

func main() {

	http.HandleFunc("/struct", AddHandler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8081", nil))
}

func AddHandler(w http.ResponseWriter, r *http.Request) {

	//parameters
	first, second := r.FormValue("first"), r.FormValue("second")
	one, err := strconv.Atoi(first)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	two, err := strconv.Atoi(second)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	m := struct{ a, b int }{one, two}
	structSum := add{Sum: m.a + m.b}

	//template - OPTIMIZE
	var html bytes.Buffer
	err = templates.Execute(&html, structSum)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html.String()))
}
