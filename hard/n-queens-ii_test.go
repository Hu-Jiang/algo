package hard

import (
	"testing"
)

func TestTotalNQueues(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 0},
		{3, 0},
		{4, 2},
		{8, 92},
	}

	for i, tt := range tests {
		got := totalNQueens(tt.n)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
