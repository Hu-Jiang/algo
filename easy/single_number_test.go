package easy

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
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
	}

	for i, tt := range tests {
		got := singleNumber(tt.nums)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
