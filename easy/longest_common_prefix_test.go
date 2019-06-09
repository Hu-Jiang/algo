package easy

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"abc", "abc", "abc"}, "abc"},
		{nil, ""},
	}

	for i, tt := range tests {
		got := longestCommonPrefix(tt.strs)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
