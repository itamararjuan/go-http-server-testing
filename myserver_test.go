package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWebServerDefaultNoArgs(t *testing.T) {
	var name = "World"
	
	server := httptest.NewServer(http.HandlerFunc(ServeHTTP))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}
	var expected = fmt.Sprintf("Hello, %s!", name)
	actual, err := ioutil.ReadAll(resp.Body)
	actualAsString := strings.Trim(string(actual), "\n\t\r")

	if err != nil {
		t.Fatal(err)
	}

	if expected != actualAsString {
		t.Errorf("Expected the message '%s'\n", expected)
	}
}

func TestWebServerWithArgs(t *testing.T) {
	var name = "Itamar"
	
	server := httptest.NewServer(http.HandlerFunc(ServeHTTP))
	defer server.Close()

	address := fmt.Sprintf("%s?name=%s",server.URL, name)

	resp, err := http.Get(address)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}
	var expected = fmt.Sprintf("Hello, %s!", name)
	actual, err := ioutil.ReadAll(resp.Body)
	actualAsString := strings.Trim(string(actual), "\n\t\r")

	if err != nil {
		t.Fatal(err)
	}

	if expected != actualAsString {
		t.Errorf("Expected the message '%s'\n", expected)
	}
}