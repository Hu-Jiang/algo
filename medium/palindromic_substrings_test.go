package medium

import (
	"testing"
)

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"", 0},
		{"abc", 3},
		{"aaa", 6},
	}

	for i, tt := range tests {
		got := countSubstrings(tt.s)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
