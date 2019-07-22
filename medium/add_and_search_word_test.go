package medium

import "testing"

func TestWordDictionary(t *testing.T) {
	wd := ConstructorWordDictionary()

	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")

	if wd.Search("pad") {
		t.Fatalf("expected pad not exist")
	}
	if !wd.Search("bad") {
		t.Fatalf("expected bad exist")
	}
	if !wd.Search(".ad") {
		t.Fatalf("expected .ad exist")
	}
	if !wd.Search("b..") {
		t.Fatalf("expected b.. exist")
	}
}
