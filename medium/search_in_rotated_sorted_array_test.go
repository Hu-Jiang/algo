package medium

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{nil, 0, -1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
	}

	for i, tt := range tests {
		got := search(tt.nums, tt.target)
		if got != tt.want {
			t.Fatalf("%d: got: %d; want: %d", i, got, tt.want)
		}
	}
}
