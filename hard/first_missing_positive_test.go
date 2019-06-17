package hard

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nil, 1},
		{[]int{1, 2, 0}, 3},
		{[]int{1, 1}, 2},
		{[]int{2, 2, 2}, 1},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
	}

	for i, tt := range tests {
		got := firstMissingPositive(tt.nums)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
