package main

import "testing"

type database struct {

}

func newDB() database {
	return database{}
}

func (db *database) close() {

}

func (db *database) Put(node []byte) string {
	return ""
}

func (db *database) Get(key string) []byte {
	return []byte{}
}

func TestDBPut(t *testing.T) {
	db := newDB();
	defer db.close()

	key := db.Put([]byte("testing node"))

	data := db.Get(key)

	s, _ := Decode(data, 0)
	if str, ok := s.([]byte); ok {
		if string(str) != "testing node" {
			t.Error("Wrong value node", str)
		}
	} else {
		t.Error("Invalid return type")
	}
}