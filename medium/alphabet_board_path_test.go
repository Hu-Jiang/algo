package medium

import (
	"testing"
)

func TestAlphabetBoardPath(t *testing.T) {
	tests := []struct {
		target string
		want   string
	}{
		{"", ""},
		{"leet", "DDR!UURRR!!DDD!"},
		{"code", "RR!DDRR!UUL!R!"},
		{"zdz", "DDDDD!UUUUURRR!LLLDDDDD!"},
	}

	for i, tt := range tests {
		got := alphabetBoardPath(tt.target)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
