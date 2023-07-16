package goalexa

import (
	"io"
	"os"
	"testing"
)

///https://github.com/patst/alexa-skills-kit-for-go/blob/master/alexa/http_test.go

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

// https://developer.amazon.com/en-US/docs/alexa/custom-skills/host-a-custom-skill-as-a-web-service.html#check-request-signature
func Test_validateAlexaRequest(t *testing.T) {

	//	payload := strings.NewReader(`{
	//    "version": "1.0",
	//    "session": {
	//        "new": true,
	//        "sessionId": "amzn1.echo-api.session.32bed169-e348-4693-b3b6-b19e7065840d",
	//        "application": {
	//            "applicationId": "amzn1.ask.skill.d89b3e52-2d85-4693-a664-bcaa258929aa"
	//        },
	//        "attributes": {},
	//        "user": {
	//            "userId": "amzn1.ask.account.AE4GAC2H2PZKDFCKBOBABXRT53B6"
	//        }
	//    },
	//    "context": {
	//
	//        "System": {
	//            "application": {
	//                "applicationId": "amzn1.ask.skill.d89b3e52-2d85-4693-a664-bcaa258929aa"
	//            },
	//            "user": {
	//                "userId": "amzn1.ask.account.AE4GAC2H2PZKDFC"
	//            },
	//            "device": {
	//                "deviceId": "amzn1.ask.device.AHISQFS2N3CVSZVHC5O5SDO",
	//                "supportedInterfaces": {}
	//            },
	//            "apiEndpoint": "https://api.amazonalexa.com",
	//            "apiAccessToken": "eyJ0eXAiOiJ"
	//        }
	//    },
	//    "request": {
	//        "type": "LaunchRequest",
	//        "requestId": "amzn1.echo-api.request.eea1669b-2d25-4078-8329-acada1c0be63",
	//        "locale": "en-US",
	//        "timestamp": "2023-07-10T06:34:03Z",
	//        "shouldLinkResultBeReturned": false
	//    }
	//}`)
	//

	jsonData, err := os.ReadFile("mocks/requestEnvelope.json")
	if err != nil {
		panic(err)
	}

	jsonDataStr := string(jsonData)
	payload := strings.NewReader(jsonDataStr)

	url := "http://localhost:9095/alexa"

	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	validSignature := `jsHzkhi2zPaFXV4gnHN4foePDtv4SqmreEDqKqJc8kUX7skhOlZ03uKYeqLOHAot98tVJc9pMdi1TRMnkQ8sr/GoReO++yGi3iAYjO8/XXL1oscx1vMUzmOLmvCO/EfF3/iEpNOb3BIJEiNhT2ZIwp7EisQi3eYLDmDaklSmPWWGVQRtcSq1EoHarMW9GrUaApu2cJdAjnF1aF3yFoLiHheN4DSW0qQ14N+ndba4C+YQBn4Ds2SXCFUyEC+q/H4A7SFioAE/qR3WYIMMfKuk1iEQOQY7jAFCS8zOjCaa4sM373T4mNUAojcgdAaHxzu2smLRzQSttTXfuemCijTigg==`

	req.Header.Add("signaturecertchainurl", "https://s3.amazonaws.com/echo.api/echo-api-cert-7.pem")
	req.Header.Add("signature", validSignature)

	w := httptest.NewRecorder()
	ValidateAlexaRequest(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "VALID" {
		t.Errorf("expected VALID got %v", string(data))
	}

	//if err := validateAlexaRequest(w, req); err != nil {
	//	fmt.Println(err)
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

}
func Test_verifyCertURL(t *testing.T) {
	//https://developer.amazon.com/en-US/docs/alexa/custom-skills/host-a-custom-skill-as-a-web-service.html#check-request-signature
	primeTests := []struct {
		name     string
		UrlName  string
		expected bool
		msg      string
	}{
		{"valid", "https://s3.amazonaws.com/echo.api/echo-api-cert.pem", true, "pass valid url"},
		{"valid", "https://s3.amazonaws.com:443/echo.api/echo-api-cert.pem", true, "pass valid url with port"},
		{"valid", "https://s3.amazonaws.com/echo.api/../echo.api/echo-api-cert.pem", true, "pass valid url with dot"},
		{"invalid", "http://s3.amazonaws.com/echo.api/echo-api-cert.pem", false, "Cannot pass url With Invalid Protocol"},
		{"invalid", "https://notamazon.com/echo.api/echo-api-cert.pem", false, "Cannot pass url With Invalid HostName"},
		{"invalid", "https://s3.amazonaws.com/EcHo.aPi/echo-api-cert.pem", false, "Cannot pass url With Invalid Path"},
		{"invalid", "https://s3.amazonaws.com:563/echo.api/echo-api-cert.pem", false, "Cannot pass url With Invalid Port"},
	}
	for _, e := range primeTests {
		result := verifyCertURL(e.UrlName)
		if e.expected && !result {
			t.Errorf("%s: expected valid but got invalid", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected invalid but got valid", e.name)
		}

	}

}

//var (
//	cachedCert *x509.Certificate
//)
