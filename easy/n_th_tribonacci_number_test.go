package easy

import (
	"testing"
)

func TestTribonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 4},
		{25, 1389537},
	}

	for i, tt := range tests {
		got := tribonacci(tt.n)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
