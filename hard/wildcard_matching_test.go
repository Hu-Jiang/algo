package hard

import "testing"

func TestIsMatching(t *testing.T) {
	tests := []struct {
		str, pattern string
		want         bool
	}{
		{"aa", "a", false},
		{"aa", "*", true},
		{"cb", "?a", false},
		{"abceb", "*a*b", true},
		{"acdcb", "a*c?b", false}, {"aaaabaaaabbbbaabbbaabbaababbabbaaaababaaabbbbbbaabbbabababbaaabaabaaaaaabbaabbbbaababbababaabbbaababbbba", "*****b*aba***babaa*bbaba***a*aaba*b*aa**a*b**ba***a*a*", true},
	}

	for i, tt := range tests {
		got := isMatching(tt.str, tt.pattern)
		if got != tt.want {
			t.Fatalf("%d. got: %v; want: %v", i, got, tt.want)
		}
	}
}
