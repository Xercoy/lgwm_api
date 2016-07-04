package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	router.HandleFunc("/", Index)

	// Listen on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}
