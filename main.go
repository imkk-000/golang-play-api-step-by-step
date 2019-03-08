package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// APIHandler - implementation of interface Handler
type APIHandler struct {
	DefaultResponseMessage string `json:"responseMessage"`
}

func (APIHandler APIHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Add("Content-Type", "application/json")
	responseMessage, err := json.Marshal(APIHandler)
	if err != nil {
		log.Fatalln(err)
	}
	responseWriter.Write(responseMessage)
}

func main() {
	log.Fatalln(http.ListenAndServe(":8080", APIHandler{
		DefaultResponseMessage: "404 Not Found",
	}))
}
