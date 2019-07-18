package medium

import "testing"

func TestSearchII(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   bool
	}{
		{nil, 0, false},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, true},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, false},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{[]int{3, 1}, 1, true},
		{[]int{3, 1}, 3, true},
		{[]int{1, 3, 1, 1, 1}, 3, true},
	}

	for i, tt := range tests {
		got := searchII(tt.nums, tt.target)
		if got != tt.want {
			t.Errorf("#%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
