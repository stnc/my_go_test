package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestClientUpperCase(t *testing.T) {
	expected := "{'data': 'dummy'}"

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))

	defer svr.Close()

	c := NewClient(svr.URL)
	res, err := c.MakeRequest()

	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	res = strings.TrimSpace(res)

	if res != expected {
		t.Errorf("expected res to be %s got %s", expected, res)
	}
}
