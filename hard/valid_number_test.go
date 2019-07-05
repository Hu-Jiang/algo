package hard

import "testing"

func TestIsNumber(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"", false},
		{".", false},
		{"0", true},
		{".1", true},
		{"..2", false},
		{"3.", true},
		{" 0.1", true},
		{"abc", false},
		{"1 a", false},
		{"2e10", true},
		{" -90e3   ", true},
		{" 1e", false},
		{"e3", false},
		{" 6e-1", true},
		{" 99e2.5 ", false},
		{"53.5e93", true},
		{" --6", false},
		{"-+3", false},
		{"95a54e53", false},
		{" +0e-", false},
		{"6ee69", false},
		{" 005047e+6", true},
		{"00", true},
	}

	for _, tt := range tests {
		got := isNumber(tt.s)
		if got != tt.want {
			t.Errorf("isNumber(%s) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
