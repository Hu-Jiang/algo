package easy

import (
	"testing"
)

func TestMovesToMakeZigzag(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nil, 0},
		{[]int{1, 2, 3}, 2},
		{[]int{9, 6, 1, 6, 2}, 4},
	}

	for i, tt := range tests {
		got := movesToMakeZigzag(tt.nums)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
