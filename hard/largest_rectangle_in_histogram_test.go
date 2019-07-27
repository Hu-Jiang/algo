package hard

import (
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{nil, 0},
		{[]int{2, 2, 2}, 6},
		{[]int{1, 2, 3}, 4},
		{[]int{3, 2, 1}, 4},
		{[]int{2, 1, 5, 6, 2, 3}, 10},
	}

	for i, tt := range tests {
		got := largestRectangleArea(tt.heights)
		if got != tt.want {
			t.Errorf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
