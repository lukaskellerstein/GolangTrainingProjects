package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleStructAdd(t *testing.T) {

	r := request(t, "/?first=20&second=30")

	rw := httptest.NewRecorder()

	AddHandler(rw, r)
	if rw.Code == 500 {
		t.Fatal("Internal server Error: " + rw.Body.String())
	}
	if rw.Body.String() != "<h2>Here is the sum 50</h2>" {
		t.Fatal("Expected " + rw.Body.String())
	}

}

func request(t testing.TB, url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}
