package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Req: http://localhost:1234/upper?word=abc
// Res: ABC
func upperCaseHandler(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request")
		return
	}
	word := query.Get("word")
	if len(word) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing word")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, strings.ToUpper(word))
}

func main() {
	http.HandleFunc("/upper", upperCaseHandler)
	log.Fatal(http.ListenAndServe(":1234", nil))
}
jsHzkhi2zPaFXV4gnHN4foePDtv4SqmreEDqKqJc8kUX7skhOlZ03uKYeqLOHAot98tVJc9pMdi1TRMnkQ8sr/GoReO++yGi3iAYjO8/XXL1oscx1vMUzmOLmvCO/EfF3/iEpNOb3BIJEiNhT2ZIwp7EisQi3eYLDmDaklSmPWWGVQRtcSq1EoHarMW9GrUaApu2cJdAjnF1aF3yFoLiHheN4DSW0qQ14N+ndba4C+YQBn4Ds2SXCFUyEC+q/H4A7SFioAE/qR3WYIMMfKuk1iEQOQY7jAFCS8zOjCaa4sM373T4mNUAojcgdAaHxzu2smLRzQSttTXfuemCijTigg==