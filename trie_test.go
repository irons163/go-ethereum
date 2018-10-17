package main

import (
	"testing"
)

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

func TestTrieUpdate(t *testing.T) {
	db := NewMemDatabase()
	trie := NewTrie(db)

	trie.PutSatae([]byte("dog"), "puppy")
	trie.PutSatae([]byte("dogglesworth"), "cat")

	data := trie.Get([]byte(trie.root)) // 透過hash查找，得到root node
	data = trie.Get([]byte(DecodeNode(data)[1])) // root node 有兩個欄位，key , value，取value欄位，branch node
	data = trie.Get([]byte(DecodeNode(data)[6])) // expand node有26個欄位，對應字母a~z，取第7個欄位(index==6)，得到leaf node
	PrintSliceReal(DecodeNode(data)) // 印出leaf node中的真實data:["7lesworth","cat"]
	PrintSlice(DecodeNode(data)) // 印出易讀data：["glesworth","cat"]
}