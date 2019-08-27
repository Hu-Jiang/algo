package hard

import "testing"

func TestIsScramble(t *testing.T) {
	tests := []struct {
		s1, s2 string
		want   bool
	}{
		{"abc", "", false},
		{"abc", "def", false},
		{"great", "rgeat", true},
		{"abcde", "caebd", false},
	}

	for i, tt := range tests {
		got := isScramble(tt.s1, tt.s2)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
