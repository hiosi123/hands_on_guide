package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startTestHTTPServer() *httptest.Server {
	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Hello World")
			}))

	return ts
}

func Test_FetchRemoteResource(t *testing.T) {
	ts := startTestHTTPServer()
	defer ts.Close()
	fmt.Println(ts.URL)
	expected := "Hello World"
	data, err := fetchRemoteResource(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	if expected != string(data) {
		t.Errorf("Expected response to be: %s, Got: %s", expected, data)
	}
}
