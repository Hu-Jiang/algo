package hard

import (
	"testing"
)

func TestJump(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nil, 0},
		{[]int{0}, 0},
		{[]int{1}, 0},
		{[]int{1, 0}, 1},
		{[]int{1, 1}, 1},
		{[]int{1, 2, 3}, 2},
		{[]int{1, 3, 2}, 2},
		{[]int{3, 2, 1}, 1},
		{[]int{2, 3, 1}, 1},
		{[]int{1, 1, 1, 1}, 3},
		{[]int{2, 3, 1, 1, 4}, 2},
	}

	for i, tt := range tests {
		got := jump(tt.nums)
		if got != tt.want {
			t.Fatalf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
