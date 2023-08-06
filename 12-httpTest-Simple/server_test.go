package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestHandler(t *testing.T) {
	// expected := "Hello john"
	req := httptest.NewRequest(http.MethodGet, "/greet?name=john", nil)
	w := httptest.NewRecorder()
	RequestHandler(w, req)
	res := w.Result()

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	//empJSON, err := json.MarshalIndent(data, "", "  ")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Printf("data oku\n %s\n", string(empJSON))

	if err != nil {
		t.Errorf("Error: %v", err)
	}
	t.Errorf("Error:")
	if data != nil {
		t.Errorf("Expected Hello john but got %v", string(data))
	}

}
