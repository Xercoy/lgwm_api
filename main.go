package main

import (
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
)

var global_db *BoltDB

func main() {
	port := ":8080"
	router := NewRouter()

	global_db = NewBoltDB("lgwm.db", "db_vars", 0777, &bolt.Options{Timeout: 1 * time.Second})
	//	defer db.Close()

	if err := global_db.Open(); err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("Database opened.")

	log.Printf("Starting API Server on %s...", port[1:])

	// Listen on port.
	log.Fatal(http.ListenAndServe(port, router))
}
