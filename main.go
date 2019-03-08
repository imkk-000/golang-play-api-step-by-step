package main

import "net/http"

func APIHandler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.Write([]byte("My API!"))
}

func main() {
	http.HandleFunc("/", APIHandler)
	http.ListenAndServe(":8080", nil)
}
