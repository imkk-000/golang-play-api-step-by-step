package main_test

import (
	. "go-api"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_APIHandler(t *testing.T) {
	expectedResponseMessage := "404 Not Found"
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	responseWriter := httptest.NewRecorder()

	apiHandler := APIHandler("404 Not Found")
	apiHandler.ServeHTTP(responseWriter, request)
	actualResponseMessage, _ := ioutil.ReadAll(responseWriter.Body)

	if expectedResponseMessage != string(actualResponseMessage) {
		t.Errorf("expected '%s' but it got '%s'", expectedResponseMessage, string(actualResponseMessage))
	}
}
