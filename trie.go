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

		if len(key) >= len(k) && CompareIntSlice(k, key[:len(k)]) { // 如果 Prefix-Hex相同，遞歸剩下的hex數組。
			return t.GetState(v, key[len(k):])
		} else {
			return ""
		}
	} else if len(currentNode) == 17 {
		return t.GetState(currentNode[key[0]], key[1:])
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
		v := currentNode[1]

		// Matching key pair (ie. there's already an object with this key)
		// 如果key已經存在，更新value。
		if CompareIntSlice(k, key) {
			return string(t.Put([]string{ CompactEncode(key), value }))
		}

		matchingLength, newHash := t.MatchingNibble(key, value, k, v)

		if matchingLength == 0 {
			// End of the chain, return
			return newHash
		} else {
			newNode := []string{ CompactEncode(key[:matchingLength]), newHash }
			return string(t.Put(newNode))
		}
	} else {
		// Copy the current node over to the new node and replace the first nibble in the key
		newNode := make([]string, 17); copy(newNode, currentNode)
		newNode[key[0]] = t.InsertState(currentNode[key[0]], key[1:], value)

		return string(t.Put(newNode))
	}

	fmt.Println("Key is not exist.")
	return ""
}

func (t *Trie) MatchingNibble(key []int, value string, k []int, v string) (int, string) {
	var newHash string
	matchingLength := MatchingNibbleLength(key, k)
	if matchingLength == len(k) {
		// Insert the hash, creating a new node
		newHash = t.InsertState(v, key[matchingLength:], value)
	} else {
		// Expand the 2 length slice to a 17 length slice
		oldNode := t.InsertState("", k[matchingLength+1:], v)
		newNode := t.InsertState("", key[matchingLength+1:], value)
		// Create an expanded slice
		scaledSlice := make([]string, 17)
		// Set the copied and new node
		scaledSlice[k[matchingLength]] = oldNode
		scaledSlice[key[matchingLength]] = newNode

		newHash = string(t.Put(scaledSlice))
	}

	return matchingLength, newHash
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

func PrintSliceReal(slice []string) {
	fmt.Printf("[")
	for i, val := range slice {
		fmt.Printf("%q", val)
		if i != len(slice)-1 { fmt.Printf(",") }
	}
	fmt.Printf("]\n")
}

func PrintSlice(slice []string) {
	fmt.Printf("[")
	for i, val := range slice {
		if i == 0 {
			var valStr string
			for i, v := range val {
				if i == 0 {
					if v == ' ' || v == 0 {
						continue
					}

					index := v - '0'
					s := toCharStr(int(index))
					valStr += s
				} else {
					valStr += string(v)
				}
			}
			val = valStr
		}
		fmt.Printf("%q", val)
		if i != len(slice)-1 { fmt.Printf(",") }
	}
	fmt.Printf("]\n")
}

func toCharStr(i int) string {
	return string('a' - 1 + i)
}
