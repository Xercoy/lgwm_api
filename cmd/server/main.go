package main

import (
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/xercoy/lgwm_api"
)

func main() {
	port := ":8080"
	router := lgwm_api.NewRouter()

	lgwm_api.GlobalDB = lgwm_api.NewBoltDB("lgwm.db", "BlogPosts", 0777, &bolt.Options{Timeout: 1 * time.Second})
	//	defer db.Close()

	if err := lgwm_api.GlobalDB.Open(); err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("Database opened.")

	log.Printf("Starting API Server on %s...", port[1:])

	// Listen on port.
	log.Fatal(http.ListenAndServe(port, router))
}
