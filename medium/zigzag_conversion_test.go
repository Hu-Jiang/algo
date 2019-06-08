package medium

import (
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		s    string
		n    int
		want string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"abcdefg", 1, "abcdefg"},
		{"abcd", 2, "acbd"},
		{"", 10, ""},
	}

	for i, tt := range tests {
		got := convert(tt.s, tt.n)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
