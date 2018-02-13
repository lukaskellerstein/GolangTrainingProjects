package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockDatabase struct{}

func (MockDatabase) GetText() string {
	return "some text from Mock database"
}

func NewMockDatabase() AppDatabase {
	return MockDatabase{}
}

func TestHelloHandler(t *testing.T) {

	mockdb := NewMockDatabase()

	//create handler
	homeHandle := HomeHandler(mockdb)

	//create request
	req, _ := http.NewRequest("GET", "", nil)

	//create recording
	w := httptest.NewRecorder()

	//call the endpoint
	homeHandle.ServeHTTP(w, req)

	// RESULT

	// TEST IF server returns OK 200
	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}

	// TEST IF webpage contains text from database
	if !strings.Contains(w.Body.String(), "some text from Mock database") {
		t.Errorf("Home page didn't return the right thing : %s ", w.Body.String())
	}
}
