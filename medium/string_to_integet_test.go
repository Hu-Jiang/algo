package medium

import (
	"testing"
)

func TestAtoi(t *testing.T) {
	tests := []struct {
		str  string
		want int
	}{
		{"12345", 12345},
		{" +1", 1},
		{" ++++1", 0},
		{" - 345111aaaa", 0},
		{"4213 ooo", 4213},
		{"-91283472332", -2147483648},
		{"9223372036854775808", 2147483647},
	}

	for i, tt := range tests {
		got := myAtoi(tt.str)
		if got != tt.want {
			t.Fatalf("%d: got: %d; want: %d", i, got, tt.want)
		}
	}
}
