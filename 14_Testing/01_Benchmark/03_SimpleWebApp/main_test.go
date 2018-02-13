package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Benchmark_HomeHandler(b *testing.B) {

	for index := 0; index < b.N; index++ {

		//create handler
		homeHandle := homeHandler()

		//create request
		req, _ := http.NewRequest("GET", "", nil)

		//create recording
		w := httptest.NewRecorder()

		//call the endpoint
		homeHandle.ServeHTTP(w, req)

	}

}
