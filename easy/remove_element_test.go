package easy

import "testing"

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
		want int
	}{
		{nil, 0, 0},
		{[]int{1}, 1, 0},
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for i, tt := range tests {
		got := removeElement(tt.nums, tt.val)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
