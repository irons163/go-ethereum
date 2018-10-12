package main

import "testing"

func TestDBPut(t *testing.T) {
	db := NewDatabase()
	defer db.Close()

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