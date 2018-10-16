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
	//trie.PutSatae([]byte("dogglesworth"), "cat")
	v := trie.GetSatae([]byte("dog"))
	if string(v) != "puppy" {
		t.Error("trie PutSatae GetSatae fail")
	}

	data := trie.Get([]byte(trie.root))
	PrintSlice(DecodeNode(data))

}