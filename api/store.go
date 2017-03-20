package api

import (
	"errors"
	"github.com/boltdb/bolt"
)

const (
	BUCKET_NAME         = "urls"
	REVERSE_BUCKET_NAME = "reverseurls"
)

type Store interface {
	Add(url string) error
	Get(key string) ([]byte, error)
}

type UrlStore struct {
	db *DB
}

var UrlNotExists = errors.New("Url not exists")

// Creates a new UrlStore and initializes it with two buckets: one for the urls indexed by its key,
// and another reversed in order to do a fast reverse search.
func NewUrlStore(db *DB) (*UrlStore, error) {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BUCKET_NAME))
		_, err = tx.CreateBucketIfNotExists([]byte(REVERSE_BUCKET_NAME))
		return err
	})

	if err != nil {
		return &UrlStore{}, err
	}

	return &UrlStore{db}, nil
}

// Adds a new url to the store
func (u *UrlStore) Add(key, url string) error {
	return u.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET_NAME))
		rb := tx.Bucket([]byte(REVERSE_BUCKET_NAME))
		err := b.Put([]byte(key), []byte(url))
		err = rb.Put([]byte(url), []byte(key))
		return err
	})
}

// Gets the url from its hash key
func (u *UrlStore) Get(hash string) (string, error) {
	var value []byte
	u.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET_NAME))
		value = b.Get([]byte(hash))
		return nil
	})

	if value == nil {
		return "", UrlNotExists
	}

	return string(value), nil
}

// Gets the hash key by the url
func (u *UrlStore) ReverseGet(url string) (string, error) {
	var key []byte
	u.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(REVERSE_BUCKET_NAME))
		key = b.Get([]byte(url))
		return nil
	})

	if key == nil {
		return "", UrlNotExists
	}

	return string(key), nil
}
