package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBlackHandler(t *testing.T) {

	blackHandlerFunc := http.HandlerFunc(blackHandler)
	if r, err := http.NewRequest("GET", "", nil); err != nil {
		t.Errorf("%v", err)
	} else {
		recorder := httptest.NewRecorder()
		blackHandlerFunc.ServeHTTP(recorder, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}
}
