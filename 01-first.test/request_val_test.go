package goalexa

import (
	"testing"
)

func Test_verifyCertURL(t *testing.T) {

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
