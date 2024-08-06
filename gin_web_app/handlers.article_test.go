package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test that a GET request to the home page returns the home page with the HTTP status code 200 for unauthenticated users

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	// Create a response recorder

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
    // Test that the http status code is 200
    statusOK := w.Code == http.StatusOK

    // Test that the page title is "Home Page"
    // You can carry out a lot more detailed tests using libraries that can
    // parse and process HTML pages
    p, err := ioutil.ReadAll(w.Body)
    pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

    return statusOK && pageOK
  })
}
