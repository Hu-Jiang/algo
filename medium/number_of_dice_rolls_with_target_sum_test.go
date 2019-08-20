package medium

import "testing"

func TestNumRollsToTarget(t *testing.T) {
	tests := []struct {
		d, f, target int
		want         int
	}{
		{1, 6, -1, 0},
		{1, 6, 10, 0},
		{1, 6, 3, 1},
		{2, 6, 1, 0},
		{2, 6, 7, 6},
		{2, 5, 10, 1},
		{1, 2, 3, 0},
		{30, 30, 500, 222616187},
	}

	for i, tt := range tests {
		got := numRollsToTarget(tt.d, tt.f, tt.target)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %d", i, got, tt.want)
		}
	}
}
