package main

import (
	"encoding/hex"
	"testing"
)

var testsource = `{"Inputs":{
     "doe": "reindeer",
     "dog": "puppy",
     "dogglesworth": "cat"
   },
   "Expectation":"e378927bfc1bd4f01a2e8d9f59bd18db8a208bb493ac0b00f93ce51d4d2af76c"
 }`

type TestSource struct {
	Inputs map[string]string
	Expectation string
}

type TestRunner struct {
	source *TestSource
}

func NewTestRunner() TestRunner {
	testRunner := TestRunner{}
	testRunner.source = &TestSource{map[string]string{"A":"B"}, ""}
	return testRunner
}

func (t *TestRunner) RunFromString(source string, putToTrie func(*TestSource)) {
	putToTrie(t.source)
}

func TestTestRunner(t *testing.T) {
	db := NewMemDatabase()
	trie := NewTrie(db)

	runner := NewTestRunner()
	runner.RunFromString(testsource, func(source *TestSource) {
		for key, value := range source.Inputs {
			trie.PutSatae([]byte(key), value)
		}

		if hex.EncodeToString([]byte(trie.root)) != source.Expectation {
			t.Error("trie root did not match")
		}
	})
}
