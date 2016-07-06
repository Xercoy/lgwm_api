package main

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/boltdb/bolt"

	lgwm "github.com/xercoy/lgwm-api"
	reap "github.com/xercoy/reaper"
)

func main() {
	databaseName := "lgwm.db"
	bucketName := "BlogPosts"

	var err error

	db := lgwm.NewBoltDB(databaseName, bucketName, 0777, &bolt.Options{Timeout: 1 * time.Second})

	err = db.Open()
	defer db.Close()
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	log.Printf("Database opened...")

	log.Printf("Retrieving json blog posts from stdin...")

	gi, err := reap.Import(os.Stdin)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	posts, err := gi.GetPosts()
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	for id, p := range posts {
		log.Printf("Importing Post %d, '%s'", p.ID, p.Title)

		postAsBytes, err := json.Marshal(p)
		if err != nil {
			log.Fatalf("%v\n", err)
		}

		err = AddPost(db, bucketName, strconv.Itoa(id), postAsBytes)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
	}

	log.Printf("Import Completed.")
}

func AddPost(db *lgwm.BoltDB, bucketName string, key string, value []byte) error {
	return db.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		id, _ := bucket.NextSequence()

		return bucket.Put(itob(int(id)), value)
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
