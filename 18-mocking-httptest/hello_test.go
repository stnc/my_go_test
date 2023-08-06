package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(HelloWorld))
	defer testServer.Close()
	testClient := testServer.Client()

	resp, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Errorf("Get error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("response code is not 200: %d", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("io.ReadAll error: %v", err)
	}
	if string(data) != "Hello World\n" {
		t.Error("response body does not equal to Hello World")
	}
}
