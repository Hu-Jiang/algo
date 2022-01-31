package hard

import (
	"testing"
)

func TestFindWords(t *testing.T) {
	tests := []struct {
		board [][]byte
		words []string
		want  []string
	}{
		{
			board: nil,
			words: []string{"abc", "def"},
			want:  nil,
		},
		{
			board: [][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			},
			words: []string{"oath", "pea", "eat", "rain"},
			want:  []string{"oath", "eat"},
		},
		{
			board: [][]byte{
				{'a', 'b'},
				{'a', 'a'},
			},
			words: []string{"aba", "baa", "bab", "aaab", "aaa", "aaaa", "aaba"},
			want:  []string{"aaa", "aaab", "aba", "baa", "aaba"},
		},
	}

	for i, tt := range tests {
		got := findWords(tt.board, tt.words)
		assertResult(t, i, got, tt.want)
	}
}

func assertResult(t *testing.T, i int, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("#%d. length not equal. len(got) = %d, len(want) = %d", i, len(got), len(want))
	}

	findAll := func(a, b []string) bool {
		for _, va := range a {
			find := false
			for _, vb := range b {
				if va == vb {
					find = true
					break
				}
			}
			if !find {
				return false
			}
		}
		return true
	}

	if !findAll(got, want) || !findAll(want, got) {
		t.Fatalf("#%d. content not equal. got %v, want %v", i, got, want)
	}
}
