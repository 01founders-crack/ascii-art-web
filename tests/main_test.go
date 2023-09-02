// tests/main_test.go

package main_test

import (
	"ascii-art-web/handlers"
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeTemplate(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ServeTemplate)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("ServeTemplate returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	// You can add more assertions to test the content of the response if needed
}

func TestHandleAsciiArt(t *testing.T) {
	// Construct a sample POST request with form data
	form := []byte("artstyle=Standard&text=Hello%2C+ASCII+Art%21")
	req, err := http.NewRequest("POST", "/ascii-art", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Body = ioutil.NopCloser(bytes.NewReader(form))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleAsciiArt)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("HandleAsciiArt returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	// You can add more assertions to test the content of the response if needed
}
