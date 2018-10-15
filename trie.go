package main

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
	node, _ := t.db.Get(key)
	return node
}

func (t *Trie) Put(node []byte) []byte {
	enc := Encode(node)
	sha := Sha256Bin(enc)

	t.db.Put([]byte(sha), enc)

	return sha
}
