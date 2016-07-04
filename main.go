package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	router := NewRouter()

	log.Printf("Starting API Server on %s...", port[1:])

	// Listen on port.
	log.Fatal(http.ListenAndServe(port, router))
}
