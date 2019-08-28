package medium

import (
	"testing"
)

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"", 0},
		{"01", 0},
		{"10", 1},
		{"101", 1},
		{"909", 0},
		{"12", 2},
		{"226", 3},
		{"8256", 2},
	}

	for i, tt := range tests {
		got := numDecodings(tt.s)
		if got != tt.want {
			t.Errorf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
