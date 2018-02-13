package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkHandleStructAdd(b *testing.B) {
	r := request(b, "/?first=20&second=30")
	for i := 0; i < b.N; i++ {
		rw := httptest.NewRecorder()
		AddHandler(rw, r)
	}

}

func request(t testing.TB, url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}
