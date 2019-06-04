package medium

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for i, tt := range tests {
		got := lengthOfLongestSubstring(tt.s)
		if got != tt.want {
			t.Fatalf("case %d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
