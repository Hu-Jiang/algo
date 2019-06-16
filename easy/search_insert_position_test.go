package easy

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{nil, 0, 0},
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
	}

	for i, tt := range tests {
		got := searchInsert(tt.nums, tt.target)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
