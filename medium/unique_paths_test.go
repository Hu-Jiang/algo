package medium

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m, n int
		want int
	}{
		{0, 0, 0},
		{1, 1, 1},
		{3, 2, 3},
		{7, 3, 28},
	}

	for i, tt := range tests {
		got := uniquePaths(tt.m, tt.n)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
