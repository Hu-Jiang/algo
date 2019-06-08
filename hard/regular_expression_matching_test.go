package hard

import "testing"

func TestIsMatch(t *testing.T) {
	tests := []struct {
		text    string
		pattern string
		want    bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
	}

	for i, tt := range tests {
		got := isMatch(tt.text, tt.pattern)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
