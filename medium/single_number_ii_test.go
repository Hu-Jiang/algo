package medium

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: nil,
			want: 0,
		},
		{
			nums: []int{2, 2, 3, 2},
			want: 3,
		},
		{
			nums: []int{0, 1, 0, 1, 0, 1, 99},
			want: 99,
		},
		{
			nums: []int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2},
			want: -4,
		},
	}

	for i, tt := range tests {
		got := singleNumber(tt.nums)
		if got != tt.want {
			t.Errorf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
