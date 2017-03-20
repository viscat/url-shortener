package api

import (
	"testing"
	"github.com/boltdb/bolt"
	"io/ioutil"
	"os"
)



func openTestDB() (*DB, error) {

	// Retrieve a temporary path.
	f, err := ioutil.TempFile("", "")
	if err != nil {
		panic("temp file: " + err.Error())
	}
	path := f.Name()
	f.Close()
	os.Remove(path)
	// Open the database.
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		panic("open: " + err.Error())
	}
	// Return wrapped type.
	return &DB{db}, nil
}

// Close and delete Bolt database.
func (db *DB) Close() {
	defer os.Remove(db.Path())
	db.DB.Close()
}

func TestNewUrlStore(t *testing.T) {
	db, err := openTestDB()
	defer db.Close()
	if err != nil {
		panic(err)
	}
	store, _ := NewUrlStore(db)

	store.Add("a", "b")
	value, _ := store.Get("a")
	if value != "b" {
		t.Fatal("Nope")
	}
}