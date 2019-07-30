package easy

import "testing"

func TestArrayPairSum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nil, 0},
		{[]int{1, 2, 3, 4}, 4},
		{[]int{6, 5, 4, 3, 2, 1}, 9},
	}

	for i, tt := range tests {
		got := arrayPairSum(tt.nums)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
