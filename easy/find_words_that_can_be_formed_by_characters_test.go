package easy

import (
	"testing"
)

func TestCountCharacters(t *testing.T) {
	tests := []struct {
		words []string
		chars string
		want  int
	}{
		{nil, "", 0},
		{nil, "abc", 0},
		{
			[]string{"abc"},
			"",
			0,
		},
		{
			[]string{"cat", "bt", "hat", "tree"},
			"atach",
			6,
		},
		{
			[]string{"hello", "world", "leetcode"},
			"welldonehoneyr",
			10,
		},
	}

	for i, tt := range tests {
		got := countCharacters(tt.words, tt.chars)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}

}
