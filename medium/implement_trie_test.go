package medium

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("word")
	if !trie.Search("word") {
		t.Fatalf("excepted word exist")
	}
	if trie.Search("wor") {
		t.Fatalf("excepted wor not is a word in trie")
	}
	if !trie.StartsWith("wor") {
		t.Fatalf("excepted wor is a prefix in trie")
	}

	if trie.Search("apple") {
		t.Fatalf("excepted apple not exist")
	}
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Fatalf("excepted apple exist")
	}
	if trie.Search("app") {
		t.Fatalf("execpted app not exist")
	}
	if !trie.StartsWith("app") {
		t.Fatalf("excepted app is a prefix in trie")
	}
	trie.Insert("app")
	if !trie.Search("app") {
		t.Fatalf("excepted app is a word in trie")
	}
}
