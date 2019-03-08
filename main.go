package main

import "net/http"

// APIHandler - implementation of interface Handler
type APIHandler struct{}

func (APIHandler APIHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.Write([]byte("My API!"))
}

func main() {
	http.Handle("/", APIHandler{})
	http.ListenAndServe(":8080", nil)
}
