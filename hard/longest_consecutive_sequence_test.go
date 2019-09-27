package hard

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: nil,
			want: 0,
		},
		{
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
	}

	for i, tt := range tests {
		got := longestConsecutive(tt.nums)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
