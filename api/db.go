package api

import (
	"github.com/boltdb/bolt"
	"time"
)

type DB struct {
	*bolt.DB
}

func OpenDB() (*DB, error) {
	db, err := bolt.Open("urls.db", 0600, &bolt.Options{Timeout: time.Duration(1 * time.Second)})
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
