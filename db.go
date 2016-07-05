package main

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"os"

	"github.com/boltdb/bolt"
)

type BoltDB struct {
	DB      *bolt.DB
	Path    string
	Mode    os.FileMode
	Options *bolt.Options
	Bucket  string
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

	b.DB, err = bolt.Open(b.Path, b.Mode, b.Options)
	if err != nil {
		return err
	}

	//b.DB = db

	err = b.PrepareDB()

	if b.DB == nil {
		log.Printf("DB is nil")
	} else {
		log.Printf("DB is not nil: %v", b)
	}

	if err != nil || b.DB == nil {
		return err
	}
	return err
}

func (b *BoltDB) PrepareDB() error {
	var err error

	// Create bucket if it doesn't exist.
	err = b.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(b.Bucket))
		if err != nil {
			//			return fmt.Errorf("Bucket read/create error: %s", err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func GetAllPosts(db *BoltDB) ([]Post, error) {
	var err error

	var posts []Post

	err = db.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(db.Bucket))

		cursor := bucket.Cursor()

		for k, value := cursor.First(); k != nil; k, value = cursor.Next() {
			var p Post

			err := json.Unmarshal(value, &p)
			if err != nil {
				return err
			}

			posts = append(posts, p)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	log.Printf("POSTS=%v", posts)

	return posts, nil
}

func AddPost(db *BoltDB, bucket string, newPost Post) error {
	err := db.DB.Update(func(tx *bolt.Tx) error {

		log.Println(db.Bucket)

		bucket := tx.Bucket([]byte(db.Bucket))

		id, _ := bucket.NextSequence()
		newPost.ID = int(id)

		// Marshal data into bytes.
		buf, err := json.Marshal(newPost)
		if err != nil {
			return err
		}

		return bucket.Put(itob(newPost.ID), buf)
	})
	if err != nil {
		log.Fatalf("%v", err)
	}

	return nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
