package main

import (
	"os"

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
}

func NewBoltDB(path string, mode os.FileMode, options *bolt.Options) *BoltDB {
	db := new(BoltDB)
	db.Path = path
	db.Mode = mode
	db.Options = options

	return db
}

func (b *BoltDB) Close() error {
	return b.DB.Close()
}

func (b *BoltDB) Open() error {
	db, err := bolt.Open(b.Path, b.Mode, b.Options)
	if err != nil {
		return err
	}

	b.DB = db

	return nil
}
