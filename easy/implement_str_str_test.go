package easy

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"", "", 0},
		{"", "a", -1},
		{"a", "", 0},
		{"aa", "aaa", -1},
		{"aaa", "aa", 0},
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
	}

	for i, tt := range tests {
		got := strStr(tt.haystack, tt.needle)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
