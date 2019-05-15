package main

import (
	"fmt"

	"github.com/langgo/bolt"
)

func main() {
	db, err := bolt.Open("db.bin", 0666, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(db.String())

	if err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("fanzhaobing"))
		if err != nil {
			return err
		}

		return bucket.Put([]byte("key"), []byte("value"))
	}); err != nil {
		panic(err)
	}
}
