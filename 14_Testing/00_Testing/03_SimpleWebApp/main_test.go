package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HomeHandler(t *testing.T) {

	//create handler
	homeHandle := homeHandler()

	//create request
	req, _ := http.NewRequest("GET", "", nil)

	//create recording
	w := httptest.NewRecorder()

	//call the endpoint
	homeHandle.ServeHTTP(w, req)

	//RESULT
	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}
