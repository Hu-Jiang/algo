package easy

import (
	"testing"
)

func TestIsPowerOfThree(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{-1, false},
		{0, false},
		{1, true},
		{9, true},
		{27, true},
		{45, false},
	}

	for i, tt := range tests {
		got := isPowerOfThree(tt.n)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
