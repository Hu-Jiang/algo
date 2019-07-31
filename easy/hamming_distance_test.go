package easy

import "testing"

func TestHammingDistance(t *testing.T) {
	tests := []struct {
		x, y int
		want int
	}{
		{0, 0, 0},
		{1, 1, 0},
		{1, 0, 1},
		{1, 4, 2},
	}

	for i, tt := range tests {
		got := hammingDistance(tt.x, tt.y)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
