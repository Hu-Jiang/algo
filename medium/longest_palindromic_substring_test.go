package medium

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"aaaaa", "aaaaa"},
		{"", ""},
	}

	for i, tt := range tests {
		got := longestPalindrome(tt.s)
		if got != tt.want {
			t.Fatalf("%d: got: %s; want: %s", i, got, tt.want)
		}
	}
}
