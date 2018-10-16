package main

import "fmt"

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

func (t *Trie) Put(node interface{}) []byte {
	enc := Encode(node)
	sha := Sha256Bin(enc)

	t.db.Put([]byte(sha), enc)
	return sha
}

func (t *Trie) GetSatae(key []byte) ([]byte) {
	k := CompactHexDecode(string(key))

	return []byte(t.GetState(t.root, k))
}

func (t *Trie) PutSatae(key []byte, value string) []byte {
	k := CompactHexDecode(string(key))

	t.root = t.InsertState(t.root, k, value)
	return []byte(t.root)
}

// Returns the state of an object
func (t *Trie) GetState(node string, key []int) string {
	// Return the node if key is empty (= found)
	if len(key) == 0 || node == "" {
		return node
	}

	// Fetch the encoded node from the db
	n, err := t.db.Get([]byte(node))
	if err != nil { fmt.Println("Error in GetState for node", node, "with key", key); return "" }

	// Decode it
	currentNode := DecodeNode(n)

	if len(currentNode) == 0 {
		return ""
	} else if len(currentNode) == 2 {
		// Decode the key
		k := CompactDecode(currentNode[0])
		v := currentNode[1]

		if len(key) == len(k) && CompareIntSlice(k, key[:len(k)]) {
			return v
		}
	}

	// It shouldn't come this far
	fmt.Println("GetState unexpected return")
	return ""
}

func (t *Trie) InsertState(node string, key []int, value string) string {
	if len(key) == 0 {
		return value
	}

	// Root node!
	if node == "" {
		newNode := []string{ CompactEncode(key), value }

		return string(t.Put(newNode))
	}

	// Fetch the encoded node from the db
	n, err := t.db.Get([]byte(node))
	if err != nil { fmt.Println("Error InsertState", err); return "" }

	// Decode it
	currentNode := DecodeNode(n)
	// Check for "special" 2 slice type node
	if len(currentNode) == 2 {
		// Decode the key
		k := CompactDecode(currentNode[0])
		//v := currentNode[1]

		// Matching key pair (ie. there's already an object with this key)
		// 如果key已經存在，更新value。
		if CompareIntSlice(k, key) {
			return string(t.Put([]string{ CompactEncode(key), value }))
		}
	}

	fmt.Println("Key is not exist.")
	return ""
}

func DecodeNode(data []byte) []string {
	dec, _ := Decode(data, 0)
	if slice, ok := dec.([]interface{}); ok {
		strSlice := make([]string, len(slice))

		for i, s := range slice {
			if str, ok := s.([]byte); ok {
				strSlice[i] = string(str)
			}
		}

		return strSlice
	}

	return nil
}

func PrintSlice(slice []string) {
	fmt.Printf("[")
	for i, val := range slice {
		fmt.Printf("%q", val)
		if i != len(slice)-1 { fmt.Printf(",") }
	}
	fmt.Printf("]\n")
}
