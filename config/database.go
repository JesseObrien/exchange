package config

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func DatabaseInit(cfg *Config) *bolt.DB {
	// Open the bolt db instance and
	db, err := bolt.Open(cfg.DbFilename, 0600, &bolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	// Make sure the default buckets exist
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Companies"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		_, err = tx.CreateBucketIfNotExists([]byte("Transactions"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		return nil
	})

	return db
}
