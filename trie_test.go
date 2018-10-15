package main

import (
	"testing"
)

type MemDatabase struct {
	Database

	db      map[string][]byte
}

func NewMemDatabase() (*MemDatabase) {
	db := &MemDatabase{db: make(map[string][]byte)}
	return db
}

func (db *MemDatabase) Put(key []byte, value []byte) {
	db.db[string(key)] = value
}

func (db *MemDatabase) Get(key []byte) ([]byte, error) {
	return db.db[string(key)], nil
}

// Database interface
type Database interface {
	Put(key []byte, value []byte)
	Get(key []byte) ([]byte, error)
}

type Trie struct {
	root       string
	db         Database
}

func NewTrie(db Database) *Trie {
	return &Trie{db: db, root: ""}
}

func (t *Trie) Get(key []byte) ([]byte) {
	return nil
}

func (t *Trie) Put(value []byte) []byte {
	return nil
}

func TestTriePut(t *testing.T) {
	db := NewMemDatabase()
	trie := NewTrie(db)

	key := trie.Put([]byte("testing node"))

	data := trie.Get(key)

	s, _ := Decode(data, 0)
	if str, ok := s.([]byte); ok {
		if string(str) != "testing node" {
			t.Error("Wrong value node", str)
		}
	} else {
		t.Error("Invalid return type")
	}
}