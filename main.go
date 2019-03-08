package main

import (
	"fmt"
	"log"
	"net/http"
)

// APIHandler - implementation of interface Handler
type APIHandler struct {
	DefaultResponseMessage string
}

func (APIHandler APIHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Add("Content-Type", "application/json")
	responeMessage := fmt.Sprintf(`{"responseMessage":"%s"}`, APIHandler.DefaultResponseMessage)
	responseWriter.Write([]byte(responeMessage))
}

func main() {
	log.Fatalln(http.ListenAndServe(":8080", APIHandler{
		DefaultResponseMessage: "404 Not Found",
	}))
}
