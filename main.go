package main

import (
	"log"
	"net/http"
)

// APIHandler - implementation of interface Handler
type APIHandler string

func (responseMessage APIHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte(responseMessage))
}

func main() {
	log.Fatalln(http.ListenAndServe(":8080", APIHandler("404 Not Found")))
}
