package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/boltdb/bolt"
)

type Database interface {
	Open() error
	Close() error
}

type BoltDB struct {
	DB      *bolt.DB
	Path    string
	Mode    os.FileMode
	Options *bolt.Options
	Bucket  string
	IDCount int
}

func NewBoltDB(path, bucket string, mode os.FileMode, options *bolt.Options) *BoltDB {
	db := new(BoltDB)
	db.Path = path
	db.Mode = mode
	db.Options = options
	db.Bucket = bucket

	return db
}

func (b *BoltDB) Close() error {
	return b.DB.Close()
}

func (b *BoltDB) Open() error {
	var err error

	db, err := bolt.Open(b.Path, b.Mode, b.Options)
	if err != nil {
		return err
	}

	// Create bucket if it doesn't exist.
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(b.Bucket))
		if err != nil {
			return fmt.Errorf("Bucket read/create error: %s", err)
		}

		return nil
	})

	// Get/Set DB specific statistics and variables.
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(b.Bucket))
		value := bucket.Get([]byte("id_count"))

		//		log.Println("Database ID empty. Starting at 0.", value)

		if value == nil {
			b.IDCount = 0
			log.Println("Database ID empty. Starting at 0.", value)
		} else {
			// Case the byte slice to a string, conver to integer.
			valueAsInt, err := strconv.Atoi((string)(value))
			if err != nil {
				return err
			}

			b.IDCount = valueAsInt
			log.Println("Database ID count is %d.", value)
		}

		return nil
	})

	if err != nil {
		return err
	}

	b.DB = db

	return err
}
