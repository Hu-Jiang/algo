package easy

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nil, 0},
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	for i, tt := range tests {
		got := maxSubArray(tt.nums)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
