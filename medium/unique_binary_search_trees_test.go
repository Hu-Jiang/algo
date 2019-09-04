package medium

import (
	"testing"
)

func TestNumTrees(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 5},
	}

	for i, tt := range tests {
		got := numTrees(tt.n)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
