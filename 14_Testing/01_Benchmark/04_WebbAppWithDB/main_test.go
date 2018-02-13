package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockDatabase struct{}

func (MockDatabase) GetText() string {
	return "some text from Mock database"
}

func NewMockDatabase() AppDatabase {
	return MockDatabase{}
}

func Benchmark_HelloHandler(b *testing.B) {

	for index := 0; index < b.N; index++ {

		mockdb := NewMockDatabase()

		//create handler
		homeHandle := HomeHandler(mockdb)

		//create request
		req, _ := http.NewRequest("GET", "", nil)

		//create recording
		w := httptest.NewRecorder()

		//call the endpoint
		homeHandle.ServeHTTP(w, req)

	}
}
