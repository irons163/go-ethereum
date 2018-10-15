package main

import (
	"fmt"
	"os/user"
	"path"
	"github.com/syndtr/goleveldb/leveldb"
)

type LDBDatabase struct {
	db        *leveldb.DB
}

func NewDatabase() (*LDBDatabase) {
	// This will eventually have to be something like a resource folder.
	// it works on my system for now. Probably won't work on Windows
	usr, _ := user.Current()
	dbPath := path.Join(usr.HomeDir, ".ethereum", "database")

	// Open the db
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return nil
	}

	database := &LDBDatabase{db: db}

	return database
}

func (db *LDBDatabase) Put(value []byte) string {
	key := []byte("1")
	enc := Encode(value)
	err := db.db.Put(key, enc, nil)
	if err != nil {
		fmt.Println("Error put", err)
	}
	return string(key)
}

func (db *LDBDatabase) Get(key string) []byte {
	value, err := db.db.Get([]byte(key), nil)
	if err != nil {
		fmt.Println("Error put", err)
	}
	return value
}

func (db *LDBDatabase) Close() {
	// Close the leveldb database
	db.db.Close()
}