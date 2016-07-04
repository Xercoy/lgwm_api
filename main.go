package main

import (
	"log"
	"net/http"
)

var db Database

func main() {
	port := ":8080"
	router := NewRouter()

	db = NewBoltDB("lgwm.db", 0666, nil)
	defer db.Close()

	if err := db.Open(); err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("Database opened.")

	log.Printf("Starting API Server on %s...", port[1:])

	// Listen on port.
	log.Fatal(http.ListenAndServe(port, router))
}
