package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Benchmark_HomeHandler(b *testing.B) {

	//create request
	req, _ := http.NewRequest("GET", "", nil)

	for index := 0; index < b.N; index++ {

		//create recording
		w := httptest.NewRecorder()

		//call the endpoint
		homeHandleFunc(w, req)

	}

}
