package medium

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-3, -2, -5, 3, -4}, -1, -2},
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{-1, 2, 3}, 100, 4},
		{nil, 100, 0},
	}

	for i, tt := range tests {
		got := threeSumClosest(tt.nums, tt.target)
		if got != tt.want {
			t.Fatalf("%d: got: %d; want: %d", i, got, tt.want)
		}
	}
}
